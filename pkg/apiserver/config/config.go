package config

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/apiserver/registry/pkg"
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/client/openapi"
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/config/scheme"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	genericopenapi "k8s.io/apiserver/pkg/endpoints/openapi"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/server/dynamiccertificates"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/client-go/informers"
	"k8s.io/klog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	aggregatorclient "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset"
)

const (
	// selfSignedCertDir is the dir kapp-controller self signed certificates are created in.
	selfSignedCertDir = "/home/kapp-controller/kc-agg-api-selfsigned-certs"

	bindAddress = "0.0.0.0"
	bindPort    = 10349
	TokenPath   = "/token-dir"
)

type Config struct {
	codecs      serializer.CodecFactory
	genericConf *genericapiserver.Config
}

func NewConfig(aggClient aggregatorclient.Interface) (Config, error) {
	secureServing := genericoptions.NewSecureServingOptions().WithLoopback()
	authentication := genericoptions.NewDelegatingAuthenticationOptions()
	authorization := genericoptions.NewDelegatingAuthorizationOptions().WithAlwaysAllowPaths("/healthz")

	err := generateAndApplyCerts(aggClient, secureServing)
	if err != nil {
		return Config{}, err
	}

	secureServing.BindAddress = net.ParseIP(bindAddress)
	secureServing.BindPort = bindPort

	codecs := serializer.NewCodecFactory(scheme.Scheme)
	serverConfig := genericapiserver.NewConfig(codecs)

	if err := secureServing.ApplyTo(&serverConfig.SecureServing, &serverConfig.LoopbackClientConfig); err != nil {
		return Config{}, err
	}
	if err := authentication.ApplyTo(&serverConfig.Authentication, serverConfig.SecureServing, nil); err != nil {
		return Config{}, err
	}
	if err := authorization.ApplyTo(&serverConfig.Authorization); err != nil {
		return Config{}, err
	}

	// if err := os.MkdirAll(path.Dir(TokenPath), os.ModeDir); err != nil {
	// 	return nil, fmt.Errorf("error when creating dirs of token file: %v", err)
	// }

	// if err := ioutil.WriteFile(TokenPath, []byte(serverConfig.LoopbackClientConfig.BearerToken), 0600); err != nil {
	// 	return nil, fmt.Errorf("error when writing loopback access token to file: %v", err)
	// }

	serverConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(
		openapi.GetOpenAPIDefinitions,
		genericopenapi.NewDefinitionNamer(scheme.Scheme))
	serverConfig.OpenAPIConfig.Info.Title = "Kapp-controller"
	serverConfig.MinRequestTimeout = int((2 * time.Hour).Seconds())

	return Config{codecs, serverConfig}, nil
}

func (c *Config) NewServer(informerFactory informers.SharedInformerFactory) (*genericapiserver.GenericAPIServer, error) {
	server, err := c.genericConf.Complete(informerFactory).New("kapp-controller-apiserver", genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}

	// install api groups
	packagesStorage := pkg.NewREST()
	pkgGroup := genericapiserver.NewDefaultAPIGroupInfo("packages.k14s.io", scheme.Scheme, metav1.ParameterCodec, c.codecs)
	pkgv1alpha1Storage := map[string]rest.Storage{}
	pkgv1alpha1Storage["pkg"] = packagesStorage
	pkgGroup.VersionedResourcesStorageMap["v1alpha1"] = pkgv1alpha1Storage

	err = server.InstallAPIGroup(&pkgGroup)
	if err != nil {
		return nil, err
	}

	return server, nil
}

func generateAndApplyCerts(client aggregatorclient.Interface, secureServing *genericoptions.SecureServingOptionsWithLoopback) error {
	var err error
	var caContentProvider dynamiccertificates.CAContentProvider

	// Set the PairName and CertDirectory to generate the certificate files.
	secureServing.ServerCert.CertDirectory = selfSignedCertDir
	secureServing.ServerCert.PairName = "kapp-controller"

	if err := secureServing.MaybeDefaultWithSelfSignedCerts("kapp-controller", []string{"api.kapp-controller.svc"}, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	caContentProvider, err = dynamiccertificates.NewDynamicCAContentFromFile("self-signed cert", secureServing.ServerCert.CertKey.CertFile)
	if err != nil {
		return fmt.Errorf("error reading self-signed CA certificate: %v", err)
	}

	return updateAPIService(client, caContentProvider)
}

func updateAPIService(client aggregatorclient.Interface, caProvider dynamiccertificates.CAContentProvider) error {
	klog.Info("Syncing CA certificate with APIServices")
	name := "v1alpha1.packages.k14s.io"
	apiService, err := client.ApiregistrationV1().APIServices().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("error getting APIService %s: %v", name, err)
	}
	apiService.Spec.CABundle = caProvider.CurrentCABundleContent()
	if _, err := client.ApiregistrationV1().APIServices().Update(context.TODO(), apiService, metav1.UpdateOptions{}); err != nil {
		return fmt.Errorf("error updating kapp-controller CA cert of APIService %s: %v", name, err)
	}
	return nil
}

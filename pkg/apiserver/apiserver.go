package apiserver

import (
	"fmt"
	"time"

	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/apiserver/config"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	aggregatorclient "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset"
)

type APIServer struct {
	server          *genericapiserver.GenericAPIServer
	informerFactory informers.SharedInformerFactory
}

func New(clientConfig *rest.Config) (APIServer, error) {
	aggClient, err := aggregatorclient.NewForConfig(clientConfig)
	if err != nil {
		return APIServer{}, fmt.Errorf("building aggregation client: %v", err)
	}

	coreClient, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		return APIServer{}, fmt.Errorf("building core client: %v", err)
	}

	serverConfig, err := config.NewConfig(aggClient)
	if err != nil {
		return APIServer{}, fmt.Errorf("creaeting server config: %v", err)
	}

	informerFactory := informers.NewSharedInformerFactory(coreClient, 12*time.Hour)
	server, err := serverConfig.NewServer(informerFactory)
	if err != nil {
		return APIServer{}, fmt.Errorf("building server: %v", err)
	}

	return APIServer{server, informerFactory}, nil
}

func (s *APIServer) Run() <-chan struct{} {
	stopCh := make(<-chan struct{})
	go s.informerFactory.Start(stopCh)
	go s.server.PrepareRun().Run(stopCh)

	return stopCh
}

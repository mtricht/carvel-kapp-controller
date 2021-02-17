// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/ghodss/yaml"
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_PackageInstalled_FromInstalledPackage_Successfully(t *testing.T) {
	env := BuildEnv(t)
	logger := Logger{}
	kapp := Kapp{t, env.Namespace, logger}
	kubectl := Kubectl{t, env.Namespace, logger}
	sas := ServiceAccounts{env.Namespace}
	name := "instl-pkg-test"

	// contents of this bundle (k8slt/k8slt/kappctrl-e2e-repo-bundle)
	// under examples/packaging-demo
	installPkgYaml := fmt.Sprintf(`---
apiVersion: install.package.carvel.dev/v1alpha1
kind: PackageRepository
metadata:
  name: basic.test.carvel.dev
  # cluster scoped
spec:
  fetch:
    bundle:
      image: k8slt/kappctrl-e2e-repo-bundle
---
apiVersion: install.package.carvel.dev/v1alpha1
kind: InstalledPackage
metadata:
  name: %s
  namespace: %s
  annotations:
    kapp.k14s.io/change-group: kappctrl-e2e.k14s.io/installedpackages
spec:
  serviceAccountName: kappctrl-e2e-ns-sa
  packageRef:
    publicName: pkg.test.carvel.dev
    version: 1.0.0
  values:
  - secretRef:
      name: pkg-demo-values
---
apiVersion: v1
kind: Secret
metadata:
  name: pkg-demo-values
stringData:
  values.yml: |
    #@data/values
    ---
    hello_msg: "hi"
`, name, env.Namespace) + sas.ForNamespaceYAML()

	cleanUp := func() {
		// Delete App with kubectl since kapp doesn't
		// know of App that is created by kapp-controller.
		// AllowError = true since kubectl errors if it can't
		// find resource to delete.
		kubectl.RunWithOpts([]string{"delete", "apps/" + name}, RunOpts{AllowError: true})
		kapp.Run([]string{"delete", "-a", name})
	}
	cleanUp()
	defer cleanUp()

	// Create Repo, InstalledPackage, and App from YAML
	kapp.RunWithOpts([]string{"deploy", "-a", name, "-f", "-"}, RunOpts{StdinReader: strings.NewReader(installPkgYaml)})

	kubectl.Run([]string{"wait", "--for=condition=ReconcileSucceeded", "ipkg/" + name, "--timeout", "30s"})
	kubectl.Run([]string{"wait", "--for=condition=ReconcileSucceeded", "apps/" + name, "--timeout", "30s"})
	out := kubectl.Run([]string{"get", fmt.Sprintf("apps/%s", name), "-o", "yaml"})

	var cr v1alpha1.App
	err := yaml.Unmarshal([]byte(out), &cr)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s", err)
	}

	expectedStatus := v1alpha1.AppStatus{
		GenericStatus: v1alpha1.GenericStatus{
			Conditions: []v1alpha1.AppCondition{{
				Type:   v1alpha1.ReconcileSucceeded,
				Status: corev1.ConditionTrue,
			}},
			ObservedGeneration:  1,
			FriendlyDescription: "Reconcile succeeded",
		},
		Deploy: &v1alpha1.AppStatusDeploy{
			ExitCode: 0,
			Finished: true,
		},
		Fetch: &v1alpha1.AppStatusFetch{
			ExitCode: 0,
		},
		Inspect: &v1alpha1.AppStatusInspect{
			ExitCode: 0,
		},
		Template: &v1alpha1.AppStatusTemplate{
			ExitCode: 0,
		},
		ConsecutiveReconcileSuccesses: 1,
	}

	{
		// deploy
		cr.Status.Deploy.StartedAt = metav1.Time{}
		cr.Status.Deploy.UpdatedAt = metav1.Time{}
		cr.Status.Deploy.Stdout = ""

		// fetch
		cr.Status.Fetch.StartedAt = metav1.Time{}
		cr.Status.Fetch.UpdatedAt = metav1.Time{}
		cr.Status.Fetch.Stdout = ""

		// inspect
		if !strings.Contains(cr.Status.Inspect.Stdout, "5 resources\nSucceeded") {
			t.Fatalf("Expected to find 5 resources created but got:\n%s", cr.Status.Inspect.Stdout)
		}
		if !strings.Contains(cr.Status.Inspect.Stdout, "simple-app") {
			t.Fatalf("Expected to find simple-app resources created but got:\n%s", cr.Status.Inspect.Stdout)
		}
		cr.Status.Inspect.UpdatedAt = metav1.Time{}
		cr.Status.Inspect.Stdout = ""

		// template
		cr.Status.Template.UpdatedAt = metav1.Time{}
		cr.Status.Template.Stderr = ""
	}

	if !reflect.DeepEqual(expectedStatus, cr.Status) {
		t.Fatalf("\nStatus is not same:\nExpected:\n%#v\nGot:\n%#v\n", expectedStatus, cr.Status)
	}
}

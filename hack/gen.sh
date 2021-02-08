#!/bin/bash

set -e

# Note if you are not seeing generated code, most likely it's being placed into a different folder
# (e.g. Do you have GOPATH directory structure correctly named for this project?)

rm -rf pkg/client

KAPPCTRL_PKG="github.com/vmware-tanzu/carvel-kapp-controller"

echo "generating clients"
# Generate clientset and apis code with K8s codegen tools.
$GOPATH//bin/client-gen \
  --clientset-name versioned \
  --input-base "" \
  --input "${KAPPCTRL_PKG}/pkg/apis/packages/v1alpha1,${KAPPCTRL_PKG}/pkg/apis/kappctrl/v1alpha1" \
  --output-package "${KAPPCTRL_PKG}/pkg/client/clientset" \
  --go-header-file ./hack/gen-boilerplate.txt

echo "generating deepcopy"
$GOPATH/bin/deepcopy-gen \
  --input-dirs "${KAPPCTRL_PKG}/pkg/apis/packages/v1alpha1" \
  --input-dirs "${KAPPCTRL_PKG}/pkg/apis/packages" \
  --input-dirs "${KAPPCTRL_PKG}/pkg/apis/kappctrl/v1alpha1" \
  --input-dirs "${KAPPCTRL_PKG}/pkg/apis/kappctrl" \
  -O zz_generated.deepcopy \
  --go-header-file hack/gen-boilerplate.txt

echo "generating conversion"
$GOPATH/bin/conversion-gen  \
  --input-dirs "${KAPPCTRL_PKG}/pkg/apis/packages/v1alpha1,${KAPPCTRL_PKG}/pkg/apis/packages" \
  --input-dirs "${KAPPCTRL_PKG}/pkg/apis/kappctrl/v1alpha1,${KAPPCTRL_PKG}/pkg/apis/kappctrl/" \
  -O zz_generated.conversion \
  --go-header-file hack/gen-boilerplate.txt

echo "generating openapi"
$GOPATH/bin/openapi-gen  \
  --input-dirs "${KAPPCTRL_PKG}/pkg/apis/packages/v1alpha1" \
  --input-dirs "${KAPPCTRL_PKG}/pkg/apis/kappctrl/v1alpha1" \
  --input-dirs "k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/util/intstr" \
  --input-dirs "k8s.io/api/core/v1" \
  --output-package "${KAPPCTRL_PKG}/pkg/client/openapi" \
  -O zz_generated.openapi \
  --go-header-file hack/gen-boilerplate.txt


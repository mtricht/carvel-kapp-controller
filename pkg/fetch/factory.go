// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package fetch

import (
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	"k8s.io/client-go/kubernetes"
)

type Factory struct {
	coreClient kubernetes.Interface
}

func NewFactory(coreClient kubernetes.Interface) Factory {
	return Factory{coreClient}
}

func (f Factory) NewInline(opts v1alpha1.AppFetchInline, nsName string) *Inline {
	return NewInline(opts, nsName, f.coreClient)
}

// TODO: pass v1alpha1.Vendir opts here once api is exapnded
func (f Factory) NewVendir(nsName string) *Vendir {
	return NewVendir(nsName, f.coreClient)
}

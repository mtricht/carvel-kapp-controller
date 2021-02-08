// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package kappctrl

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = schema.GroupVersion{Group: "kappctrl.k14s.io", Version: runtime.APIVersionInternal}

var (
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	localSchemeBuilder.Register(func(scheme *runtime.Scheme) error {
		scheme.AddKnownTypes(SchemeGroupVersion, &App{}, &AppList{})
		scheme.AddKnownTypes(SchemeGroupVersion, &PkgRepository{}, &PkgRepositoryList{})
		scheme.AddKnownTypes(SchemeGroupVersion, &InstalledPkg{}, &InstalledPkgList{})
		return nil
	})
}

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

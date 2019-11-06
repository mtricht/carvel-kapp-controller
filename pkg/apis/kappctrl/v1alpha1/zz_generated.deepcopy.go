// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *App) DeepCopyInto(out *App) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new App.
func (in *App) DeepCopy() *App {
	if in == nil {
		return nil
	}
	out := new(App)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *App) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCondition) DeepCopyInto(out *AppCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCondition.
func (in *AppCondition) DeepCopy() *AppCondition {
	if in == nil {
		return nil
	}
	out := new(AppCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploy) DeepCopyInto(out *AppDeploy) {
	*out = *in
	if in.Kapp != nil {
		in, out := &in.Kapp, &out.Kapp
		*out = new(AppDeployKapp)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploy.
func (in *AppDeploy) DeepCopy() *AppDeploy {
	if in == nil {
		return nil
	}
	out := new(AppDeploy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeployKapp) DeepCopyInto(out *AppDeployKapp) {
	*out = *in
	if in.MapNs != nil {
		in, out := &in.MapNs, &out.MapNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RawOptions != nil {
		in, out := &in.RawOptions, &out.RawOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Delete != nil {
		in, out := &in.Delete, &out.Delete
		*out = new(AppDeployKappDelete)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeployKapp.
func (in *AppDeployKapp) DeepCopy() *AppDeployKapp {
	if in == nil {
		return nil
	}
	out := new(AppDeployKapp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeployKappDelete) DeepCopyInto(out *AppDeployKappDelete) {
	*out = *in
	if in.RawOptions != nil {
		in, out := &in.RawOptions, &out.RawOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeployKappDelete.
func (in *AppDeployKappDelete) DeepCopy() *AppDeployKappDelete {
	if in == nil {
		return nil
	}
	out := new(AppDeployKappDelete)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppFetch) DeepCopyInto(out *AppFetch) {
	*out = *in
	if in.Inline != nil {
		in, out := &in.Inline, &out.Inline
		*out = new(AppFetchInline)
		(*in).DeepCopyInto(*out)
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(AppFetchImage)
		(*in).DeepCopyInto(*out)
	}
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(AppFetchHTTP)
		(*in).DeepCopyInto(*out)
	}
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(AppFetchGit)
		(*in).DeepCopyInto(*out)
	}
	if in.HelmChart != nil {
		in, out := &in.HelmChart, &out.HelmChart
		*out = new(AppFetchHelmChart)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppFetch.
func (in *AppFetch) DeepCopy() *AppFetch {
	if in == nil {
		return nil
	}
	out := new(AppFetch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppFetchGit) DeepCopyInto(out *AppFetchGit) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(AppFetchLocalRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppFetchGit.
func (in *AppFetchGit) DeepCopy() *AppFetchGit {
	if in == nil {
		return nil
	}
	out := new(AppFetchGit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppFetchHTTP) DeepCopyInto(out *AppFetchHTTP) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(AppFetchLocalRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppFetchHTTP.
func (in *AppFetchHTTP) DeepCopy() *AppFetchHTTP {
	if in == nil {
		return nil
	}
	out := new(AppFetchHTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppFetchHelmChart) DeepCopyInto(out *AppFetchHelmChart) {
	*out = *in
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(AppFetchHelmChartRepo)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppFetchHelmChart.
func (in *AppFetchHelmChart) DeepCopy() *AppFetchHelmChart {
	if in == nil {
		return nil
	}
	out := new(AppFetchHelmChart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppFetchHelmChartRepo) DeepCopyInto(out *AppFetchHelmChartRepo) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(AppFetchLocalRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppFetchHelmChartRepo.
func (in *AppFetchHelmChartRepo) DeepCopy() *AppFetchHelmChartRepo {
	if in == nil {
		return nil
	}
	out := new(AppFetchHelmChartRepo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppFetchImage) DeepCopyInto(out *AppFetchImage) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(AppFetchLocalRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppFetchImage.
func (in *AppFetchImage) DeepCopy() *AppFetchImage {
	if in == nil {
		return nil
	}
	out := new(AppFetchImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppFetchInline) DeepCopyInto(out *AppFetchInline) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PathsFrom != nil {
		in, out := &in.PathsFrom, &out.PathsFrom
		*out = make([]AppFetchInlineSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppFetchInline.
func (in *AppFetchInline) DeepCopy() *AppFetchInline {
	if in == nil {
		return nil
	}
	out := new(AppFetchInline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppFetchInlineSource) DeepCopyInto(out *AppFetchInlineSource) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(AppFetchInlineSourceRef)
		**out = **in
	}
	if in.ConfigMapRef != nil {
		in, out := &in.ConfigMapRef, &out.ConfigMapRef
		*out = new(AppFetchInlineSourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppFetchInlineSource.
func (in *AppFetchInlineSource) DeepCopy() *AppFetchInlineSource {
	if in == nil {
		return nil
	}
	out := new(AppFetchInlineSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppFetchInlineSourceRef) DeepCopyInto(out *AppFetchInlineSourceRef) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppFetchInlineSourceRef.
func (in *AppFetchInlineSourceRef) DeepCopy() *AppFetchInlineSourceRef {
	if in == nil {
		return nil
	}
	out := new(AppFetchInlineSourceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppFetchLocalRef) DeepCopyInto(out *AppFetchLocalRef) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppFetchLocalRef.
func (in *AppFetchLocalRef) DeepCopy() *AppFetchLocalRef {
	if in == nil {
		return nil
	}
	out := new(AppFetchLocalRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppList) DeepCopyInto(out *AppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]App, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppList.
func (in *AppList) DeepCopy() *AppList {
	if in == nil {
		return nil
	}
	out := new(AppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpec) DeepCopyInto(out *AppSpec) {
	*out = *in
	if in.Fetch != nil {
		in, out := &in.Fetch, &out.Fetch
		*out = make([]AppFetch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = make([]AppTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Deploy != nil {
		in, out := &in.Deploy, &out.Deploy
		*out = make([]AppDeploy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpec.
func (in *AppSpec) DeepCopy() *AppSpec {
	if in == nil {
		return nil
	}
	out := new(AppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppStatus) DeepCopyInto(out *AppStatus) {
	*out = *in
	if in.Fetch != nil {
		in, out := &in.Fetch, &out.Fetch
		*out = new(AppStatusLastFetch)
		(*in).DeepCopyInto(*out)
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(AppStatusLastTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.Deploy != nil {
		in, out := &in.Deploy, &out.Deploy
		*out = new(AppStatusLastDeploy)
		(*in).DeepCopyInto(*out)
	}
	if in.Inspect != nil {
		in, out := &in.Inspect, &out.Inspect
		*out = new(AppStatusInspect)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AppCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppStatus.
func (in *AppStatus) DeepCopy() *AppStatus {
	if in == nil {
		return nil
	}
	out := new(AppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppStatusInspect) DeepCopyInto(out *AppStatusInspect) {
	*out = *in
	in.UpdatedAt.DeepCopyInto(&out.UpdatedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppStatusInspect.
func (in *AppStatusInspect) DeepCopy() *AppStatusInspect {
	if in == nil {
		return nil
	}
	out := new(AppStatusInspect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppStatusLastDeploy) DeepCopyInto(out *AppStatusLastDeploy) {
	*out = *in
	in.StartedAt.DeepCopyInto(&out.StartedAt)
	in.UpdatedAt.DeepCopyInto(&out.UpdatedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppStatusLastDeploy.
func (in *AppStatusLastDeploy) DeepCopy() *AppStatusLastDeploy {
	if in == nil {
		return nil
	}
	out := new(AppStatusLastDeploy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppStatusLastFetch) DeepCopyInto(out *AppStatusLastFetch) {
	*out = *in
	in.StartedAt.DeepCopyInto(&out.StartedAt)
	in.UpdatedAt.DeepCopyInto(&out.UpdatedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppStatusLastFetch.
func (in *AppStatusLastFetch) DeepCopy() *AppStatusLastFetch {
	if in == nil {
		return nil
	}
	out := new(AppStatusLastFetch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppStatusLastTemplate) DeepCopyInto(out *AppStatusLastTemplate) {
	*out = *in
	in.UpdatedAt.DeepCopyInto(&out.UpdatedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppStatusLastTemplate.
func (in *AppStatusLastTemplate) DeepCopy() *AppStatusLastTemplate {
	if in == nil {
		return nil
	}
	out := new(AppStatusLastTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppTemplate) DeepCopyInto(out *AppTemplate) {
	*out = *in
	if in.Ytt != nil {
		in, out := &in.Ytt, &out.Ytt
		*out = new(AppTemplateYtt)
		(*in).DeepCopyInto(*out)
	}
	if in.Kbld != nil {
		in, out := &in.Kbld, &out.Kbld
		*out = new(AppTemplateKbld)
		**out = **in
	}
	if in.HelmTemplate != nil {
		in, out := &in.HelmTemplate, &out.HelmTemplate
		*out = new(AppTemplateHelmTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.Kustomize != nil {
		in, out := &in.Kustomize, &out.Kustomize
		*out = new(AppTemplateKustomize)
		**out = **in
	}
	if in.Jsonnet != nil {
		in, out := &in.Jsonnet, &out.Jsonnet
		*out = new(AppTemplateJsonnet)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppTemplate.
func (in *AppTemplate) DeepCopy() *AppTemplate {
	if in == nil {
		return nil
	}
	out := new(AppTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppTemplateHelmTemplate) DeepCopyInto(out *AppTemplateHelmTemplate) {
	*out = *in
	if in.ValuesFrom != nil {
		in, out := &in.ValuesFrom, &out.ValuesFrom
		*out = make([]AppTemplateHelmTemplateValuesSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppTemplateHelmTemplate.
func (in *AppTemplateHelmTemplate) DeepCopy() *AppTemplateHelmTemplate {
	if in == nil {
		return nil
	}
	out := new(AppTemplateHelmTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppTemplateHelmTemplateValuesSource) DeepCopyInto(out *AppTemplateHelmTemplateValuesSource) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(AppTemplateHelmTemplateValuesSourceRef)
		**out = **in
	}
	if in.ConfigMapRef != nil {
		in, out := &in.ConfigMapRef, &out.ConfigMapRef
		*out = new(AppTemplateHelmTemplateValuesSourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppTemplateHelmTemplateValuesSource.
func (in *AppTemplateHelmTemplateValuesSource) DeepCopy() *AppTemplateHelmTemplateValuesSource {
	if in == nil {
		return nil
	}
	out := new(AppTemplateHelmTemplateValuesSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppTemplateHelmTemplateValuesSourceRef) DeepCopyInto(out *AppTemplateHelmTemplateValuesSourceRef) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppTemplateHelmTemplateValuesSourceRef.
func (in *AppTemplateHelmTemplateValuesSourceRef) DeepCopy() *AppTemplateHelmTemplateValuesSourceRef {
	if in == nil {
		return nil
	}
	out := new(AppTemplateHelmTemplateValuesSourceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppTemplateJsonnet) DeepCopyInto(out *AppTemplateJsonnet) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppTemplateJsonnet.
func (in *AppTemplateJsonnet) DeepCopy() *AppTemplateJsonnet {
	if in == nil {
		return nil
	}
	out := new(AppTemplateJsonnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppTemplateKbld) DeepCopyInto(out *AppTemplateKbld) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppTemplateKbld.
func (in *AppTemplateKbld) DeepCopy() *AppTemplateKbld {
	if in == nil {
		return nil
	}
	out := new(AppTemplateKbld)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppTemplateKustomize) DeepCopyInto(out *AppTemplateKustomize) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppTemplateKustomize.
func (in *AppTemplateKustomize) DeepCopy() *AppTemplateKustomize {
	if in == nil {
		return nil
	}
	out := new(AppTemplateKustomize)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppTemplateYtt) DeepCopyInto(out *AppTemplateYtt) {
	*out = *in
	if in.Inline != nil {
		in, out := &in.Inline, &out.Inline
		*out = new(AppFetchInline)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppTemplateYtt.
func (in *AppTemplateYtt) DeepCopy() *AppTemplateYtt {
	if in == nil {
		return nil
	}
	out := new(AppTemplateYtt)
	in.DeepCopyInto(out)
	return out
}

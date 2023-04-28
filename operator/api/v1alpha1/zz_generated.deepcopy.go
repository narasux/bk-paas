//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - PaaS 平台 (BlueKing - PaaS System) available.
 * Copyright (C) 2017 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 *	http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 * to the current version of the project delivered to anyone in the future.
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlConfig) DeepCopyInto(out *AccessControlConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlConfig.
func (in *AccessControlConfig) DeepCopy() *AccessControlConfig {
	if in == nil {
		return nil
	}
	out := new(AccessControlConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addressable) DeepCopyInto(out *Addressable) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addressable.
func (in *Addressable) DeepCopy() *Addressable {
	if in == nil {
		return nil
	}
	out := new(Addressable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConfig) DeepCopyInto(out *AppConfig) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]AppEnvVar, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConfig.
func (in *AppConfig) DeepCopy() *AppConfig {
	if in == nil {
		return nil
	}
	out := new(AppConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEnvOverlay) DeepCopyInto(out *AppEnvOverlay) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make([]ReplicasOverlay, len(*in))
		copy(*out, *in)
	}
	if in.EnvVariables != nil {
		in, out := &in.EnvVariables, &out.EnvVariables
		*out = make([]EnvVarOverlay, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEnvOverlay.
func (in *AppEnvOverlay) DeepCopy() *AppEnvOverlay {
	if in == nil {
		return nil
	}
	out := new(AppEnvOverlay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEnvVar) DeepCopyInto(out *AppEnvVar) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEnvVar.
func (in *AppEnvVar) DeepCopy() *AppEnvVar {
	if in == nil {
		return nil
	}
	out := new(AppEnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppHooks) DeepCopyInto(out *AppHooks) {
	*out = *in
	if in.PreRelease != nil {
		in, out := &in.PreRelease, &out.PreRelease
		*out = new(Hook)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppHooks.
func (in *AppHooks) DeepCopy() *AppHooks {
	if in == nil {
		return nil
	}
	out := new(AppHooks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpec) DeepCopyInto(out *AppSpec) {
	*out = *in
	in.Build.DeepCopyInto(&out.Build)
	if in.Processes != nil {
		in, out := &in.Processes, &out.Processes
		*out = make([]Process, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Configuration.DeepCopyInto(&out.Configuration)
	if in.Hooks != nil {
		in, out := &in.Hooks, &out.Hooks
		*out = new(AppHooks)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvOverlay != nil {
		in, out := &in.EnvOverlay, &out.EnvOverlay
		*out = new(AppEnvOverlay)
		(*in).DeepCopyInto(*out)
	}
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
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]Addressable, len(*in))
		copy(*out, *in)
	}
	if in.HookStatuses != nil {
		in, out := &in.HookStatuses, &out.HookStatuses
		*out = make([]HookStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(Revision)
		**out = **in
	}
	if in.LastUpdate != nil {
		in, out := &in.LastUpdate, &out.LastUpdate
		*out = (*in).DeepCopy()
	}
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
func (in *BkApp) DeepCopyInto(out *BkApp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BkApp.
func (in *BkApp) DeepCopy() *BkApp {
	if in == nil {
		return nil
	}
	out := new(BkApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BkApp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BkAppList) DeepCopyInto(out *BkAppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BkApp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BkAppList.
func (in *BkAppList) DeepCopy() *BkAppList {
	if in == nil {
		return nil
	}
	out := new(BkAppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BkAppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildConfig) DeepCopyInto(out *BuildConfig) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildConfig.
func (in *BuildConfig) DeepCopy() *BuildConfig {
	if in == nil {
		return nil
	}
	out := new(BuildConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Domain) DeepCopyInto(out *Domain) {
	*out = *in
	if in.PathPrefixList != nil {
		in, out := &in.PathPrefixList, &out.PathPrefixList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Domain.
func (in *Domain) DeepCopy() *Domain {
	if in == nil {
		return nil
	}
	out := new(Domain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainGroup) DeepCopyInto(out *DomainGroup) {
	*out = *in
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]Domain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainGroup.
func (in *DomainGroup) DeepCopy() *DomainGroup {
	if in == nil {
		return nil
	}
	out := new(DomainGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainGroupMapping) DeepCopyInto(out *DomainGroupMapping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainGroupMapping.
func (in *DomainGroupMapping) DeepCopy() *DomainGroupMapping {
	if in == nil {
		return nil
	}
	out := new(DomainGroupMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainGroupMapping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainGroupMappingList) DeepCopyInto(out *DomainGroupMappingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DomainGroupMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainGroupMappingList.
func (in *DomainGroupMappingList) DeepCopy() *DomainGroupMappingList {
	if in == nil {
		return nil
	}
	out := new(DomainGroupMappingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainGroupMappingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainGroupMappingSpec) DeepCopyInto(out *DomainGroupMappingSpec) {
	*out = *in
	out.Ref = in.Ref
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]DomainGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainGroupMappingSpec.
func (in *DomainGroupMappingSpec) DeepCopy() *DomainGroupMappingSpec {
	if in == nil {
		return nil
	}
	out := new(DomainGroupMappingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainGroupMappingStatus) DeepCopyInto(out *DomainGroupMappingStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainGroupMappingStatus.
func (in *DomainGroupMappingStatus) DeepCopy() *DomainGroupMappingStatus {
	if in == nil {
		return nil
	}
	out := new(DomainGroupMappingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVarOverlay) DeepCopyInto(out *EnvVarOverlay) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVarOverlay.
func (in *EnvVarOverlay) DeepCopy() *EnvVarOverlay {
	if in == nil {
		return nil
	}
	out := new(EnvVarOverlay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hook) DeepCopyInto(out *Hook) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hook.
func (in *Hook) DeepCopy() *Hook {
	if in == nil {
		return nil
	}
	out := new(Hook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HookStatus) DeepCopyInto(out *HookStatus) {
	*out = *in
	if in.Started != nil {
		in, out := &in.Started, &out.Started
		*out = new(bool)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HookStatus.
func (in *HookStatus) DeepCopy() *HookStatus {
	if in == nil {
		return nil
	}
	out := new(HookStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressPluginConfig) DeepCopyInto(out *IngressPluginConfig) {
	*out = *in
	if in.AccessControlConfig != nil {
		in, out := &in.AccessControlConfig, &out.AccessControlConfig
		*out = new(AccessControlConfig)
		**out = **in
	}
	if in.PaaSAnalysisConfig != nil {
		in, out := &in.PaaSAnalysisConfig, &out.PaaSAnalysisConfig
		*out = new(PaaSAnalysisConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressPluginConfig.
func (in *IngressPluginConfig) DeepCopy() *IngressPluginConfig {
	if in == nil {
		return nil
	}
	out := new(IngressPluginConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MappingRef) DeepCopyInto(out *MappingRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MappingRef.
func (in *MappingRef) DeepCopy() *MappingRef {
	if in == nil {
		return nil
	}
	out := new(MappingRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaaSAnalysisConfig) DeepCopyInto(out *PaaSAnalysisConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaaSAnalysisConfig.
func (in *PaaSAnalysisConfig) DeepCopy() *PaaSAnalysisConfig {
	if in == nil {
		return nil
	}
	out := new(PaaSAnalysisConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformConfig) DeepCopyInto(out *PlatformConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformConfig.
func (in *PlatformConfig) DeepCopy() *PlatformConfig {
	if in == nil {
		return nil
	}
	out := new(PlatformConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Process) DeepCopyInto(out *Process) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Process.
func (in *Process) DeepCopy() *Process {
	if in == nil {
		return nil
	}
	out := new(Process)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectConfig) DeepCopyInto(out *ProjectConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
	out.PlatformConfig = in.PlatformConfig
	in.IngressPluginConfig.DeepCopyInto(&out.IngressPluginConfig)
	out.ResLimitConfig = in.ResLimitConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectConfig.
func (in *ProjectConfig) DeepCopy() *ProjectConfig {
	if in == nil {
		return nil
	}
	out := new(ProjectConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicasOverlay) DeepCopyInto(out *ReplicasOverlay) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicasOverlay.
func (in *ReplicasOverlay) DeepCopy() *ReplicasOverlay {
	if in == nil {
		return nil
	}
	out := new(ReplicasOverlay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResLimitConfig) DeepCopyInto(out *ResLimitConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResLimitConfig.
func (in *ResLimitConfig) DeepCopy() *ResLimitConfig {
	if in == nil {
		return nil
	}
	out := new(ResLimitConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Revision) DeepCopyInto(out *Revision) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Revision.
func (in *Revision) DeepCopy() *Revision {
	if in == nil {
		return nil
	}
	out := new(Revision)
	in.DeepCopyInto(out)
	return out
}

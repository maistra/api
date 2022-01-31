//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryConnectionStatus) DeepCopyInto(out *DiscoveryConnectionStatus) {
	*out = *in
	in.LastConnected.DeepCopyInto(&out.LastConnected)
	in.LastEvent.DeepCopyInto(&out.LastEvent)
	in.LastFullSync.DeepCopyInto(&out.LastFullSync)
	in.LastDisconnect.DeepCopyInto(&out.LastDisconnect)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryConnectionStatus.
func (in *DiscoveryConnectionStatus) DeepCopy() *DiscoveryConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(DiscoveryConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryRemoteStatus) DeepCopyInto(out *DiscoveryRemoteStatus) {
	*out = *in
	in.DiscoveryConnectionStatus.DeepCopyInto(&out.DiscoveryConnectionStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryRemoteStatus.
func (in *DiscoveryRemoteStatus) DeepCopy() *DiscoveryRemoteStatus {
	if in == nil {
		return nil
	}
	out := new(DiscoveryRemoteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryWatchStatus) DeepCopyInto(out *DiscoveryWatchStatus) {
	*out = *in
	in.DiscoveryConnectionStatus.DeepCopyInto(&out.DiscoveryConnectionStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryWatchStatus.
func (in *DiscoveryWatchStatus) DeepCopy() *DiscoveryWatchStatus {
	if in == nil {
		return nil
	}
	out := new(DiscoveryWatchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportedServiceRule) DeepCopyInto(out *ExportedServiceRule) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(ServiceImportExportLabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NameSelector != nil {
		in, out := &in.NameSelector, &out.NameSelector
		*out = new(ServiceNameMapping)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportedServiceRule.
func (in *ExportedServiceRule) DeepCopy() *ExportedServiceRule {
	if in == nil {
		return nil
	}
	out := new(ExportedServiceRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportedServiceSet) DeepCopyInto(out *ExportedServiceSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportedServiceSet.
func (in *ExportedServiceSet) DeepCopy() *ExportedServiceSet {
	if in == nil {
		return nil
	}
	out := new(ExportedServiceSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExportedServiceSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportedServiceSetList) DeepCopyInto(out *ExportedServiceSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExportedServiceSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportedServiceSetList.
func (in *ExportedServiceSetList) DeepCopy() *ExportedServiceSetList {
	if in == nil {
		return nil
	}
	out := new(ExportedServiceSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExportedServiceSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportedServiceSetSpec) DeepCopyInto(out *ExportedServiceSetSpec) {
	*out = *in
	if in.ExportRules != nil {
		in, out := &in.ExportRules, &out.ExportRules
		*out = make([]ExportedServiceRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportedServiceSetSpec.
func (in *ExportedServiceSetSpec) DeepCopy() *ExportedServiceSetSpec {
	if in == nil {
		return nil
	}
	out := new(ExportedServiceSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportedServiceSetStatus) DeepCopyInto(out *ExportedServiceSetStatus) {
	*out = *in
	if in.ExportedServices != nil {
		in, out := &in.ExportedServices, &out.ExportedServices
		*out = make([]PeerServiceMapping, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportedServiceSetStatus.
func (in *ExportedServiceSetStatus) DeepCopy() *ExportedServiceSetStatus {
	if in == nil {
		return nil
	}
	out := new(ExportedServiceSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportedServiceLocality) DeepCopyInto(out *ImportedServiceLocality) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportedServiceLocality.
func (in *ImportedServiceLocality) DeepCopy() *ImportedServiceLocality {
	if in == nil {
		return nil
	}
	out := new(ImportedServiceLocality)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportedServiceRule) DeepCopyInto(out *ImportedServiceRule) {
	*out = *in
	if in.NameSelector != nil {
		in, out := &in.NameSelector, &out.NameSelector
		*out = new(ServiceNameMapping)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportedServiceRule.
func (in *ImportedServiceRule) DeepCopy() *ImportedServiceRule {
	if in == nil {
		return nil
	}
	out := new(ImportedServiceRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportedServiceSet) DeepCopyInto(out *ImportedServiceSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportedServiceSet.
func (in *ImportedServiceSet) DeepCopy() *ImportedServiceSet {
	if in == nil {
		return nil
	}
	out := new(ImportedServiceSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImportedServiceSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportedServiceSetList) DeepCopyInto(out *ImportedServiceSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImportedServiceSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportedServiceSetList.
func (in *ImportedServiceSetList) DeepCopy() *ImportedServiceSetList {
	if in == nil {
		return nil
	}
	out := new(ImportedServiceSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImportedServiceSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportedServiceSetSpec) DeepCopyInto(out *ImportedServiceSetSpec) {
	*out = *in
	if in.Locality != nil {
		in, out := &in.Locality, &out.Locality
		*out = new(ImportedServiceLocality)
		**out = **in
	}
	if in.ImportRules != nil {
		in, out := &in.ImportRules, &out.ImportRules
		*out = make([]ImportedServiceRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportedServiceSetSpec.
func (in *ImportedServiceSetSpec) DeepCopy() *ImportedServiceSetSpec {
	if in == nil {
		return nil
	}
	out := new(ImportedServiceSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportedServiceSetStatus) DeepCopyInto(out *ImportedServiceSetStatus) {
	*out = *in
	if in.ImportedServices != nil {
		in, out := &in.ImportedServices, &out.ImportedServices
		*out = make([]PeerServiceMapping, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportedServiceSetStatus.
func (in *ImportedServiceSetStatus) DeepCopy() *ImportedServiceSetStatus {
	if in == nil {
		return nil
	}
	out := new(ImportedServiceSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerDiscoveryStatus) DeepCopyInto(out *PeerDiscoveryStatus) {
	*out = *in
	if in.Remotes != nil {
		in, out := &in.Remotes, &out.Remotes
		*out = make([]DiscoveryRemoteStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Watch.DeepCopyInto(&out.Watch)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerDiscoveryStatus.
func (in *PeerDiscoveryStatus) DeepCopy() *PeerDiscoveryStatus {
	if in == nil {
		return nil
	}
	out := new(PeerDiscoveryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerServiceMapping) DeepCopyInto(out *PeerServiceMapping) {
	*out = *in
	out.LocalService = in.LocalService
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerServiceMapping.
func (in *PeerServiceMapping) DeepCopy() *PeerServiceMapping {
	if in == nil {
		return nil
	}
	out := new(PeerServiceMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodPeerDiscoveryStatus) DeepCopyInto(out *PodPeerDiscoveryStatus) {
	*out = *in
	in.PeerDiscoveryStatus.DeepCopyInto(&out.PeerDiscoveryStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodPeerDiscoveryStatus.
func (in *PodPeerDiscoveryStatus) DeepCopy() *PodPeerDiscoveryStatus {
	if in == nil {
		return nil
	}
	out := new(PodPeerDiscoveryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceImportExportLabelSelector) DeepCopyInto(out *ServiceImportExportLabelSelector) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.Aliases != nil {
		in, out := &in.Aliases, &out.Aliases
		*out = make([]ServiceNameMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceImportExportLabelSelector.
func (in *ServiceImportExportLabelSelector) DeepCopy() *ServiceImportExportLabelSelector {
	if in == nil {
		return nil
	}
	out := new(ServiceImportExportLabelSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceKey) DeepCopyInto(out *ServiceKey) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceKey.
func (in *ServiceKey) DeepCopy() *ServiceKey {
	if in == nil {
		return nil
	}
	out := new(ServiceKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMeshPeer) DeepCopyInto(out *ServiceMeshPeer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMeshPeer.
func (in *ServiceMeshPeer) DeepCopy() *ServiceMeshPeer {
	if in == nil {
		return nil
	}
	out := new(ServiceMeshPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceMeshPeer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMeshPeerDiscoveryStatus) DeepCopyInto(out *ServiceMeshPeerDiscoveryStatus) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = make([]PodPeerDiscoveryStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Inactive != nil {
		in, out := &in.Inactive, &out.Inactive
		*out = make([]PodPeerDiscoveryStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMeshPeerDiscoveryStatus.
func (in *ServiceMeshPeerDiscoveryStatus) DeepCopy() *ServiceMeshPeerDiscoveryStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceMeshPeerDiscoveryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMeshPeerGateways) DeepCopyInto(out *ServiceMeshPeerGateways) {
	*out = *in
	out.Ingress = in.Ingress
	out.Egress = in.Egress
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMeshPeerGateways.
func (in *ServiceMeshPeerGateways) DeepCopy() *ServiceMeshPeerGateways {
	if in == nil {
		return nil
	}
	out := new(ServiceMeshPeerGateways)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMeshPeerList) DeepCopyInto(out *ServiceMeshPeerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceMeshPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMeshPeerList.
func (in *ServiceMeshPeerList) DeepCopy() *ServiceMeshPeerList {
	if in == nil {
		return nil
	}
	out := new(ServiceMeshPeerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceMeshPeerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMeshPeerRemote) DeepCopyInto(out *ServiceMeshPeerRemote) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMeshPeerRemote.
func (in *ServiceMeshPeerRemote) DeepCopy() *ServiceMeshPeerRemote {
	if in == nil {
		return nil
	}
	out := new(ServiceMeshPeerRemote)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMeshPeerSecurity) DeepCopyInto(out *ServiceMeshPeerSecurity) {
	*out = *in
	in.CertificateChain.DeepCopyInto(&out.CertificateChain)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMeshPeerSecurity.
func (in *ServiceMeshPeerSecurity) DeepCopy() *ServiceMeshPeerSecurity {
	if in == nil {
		return nil
	}
	out := new(ServiceMeshPeerSecurity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMeshPeerSpec) DeepCopyInto(out *ServiceMeshPeerSpec) {
	*out = *in
	in.Remote.DeepCopyInto(&out.Remote)
	out.Gateways = in.Gateways
	in.Security.DeepCopyInto(&out.Security)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMeshPeerSpec.
func (in *ServiceMeshPeerSpec) DeepCopy() *ServiceMeshPeerSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceMeshPeerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMeshPeerStatus) DeepCopyInto(out *ServiceMeshPeerStatus) {
	*out = *in
	in.DiscoveryStatus.DeepCopyInto(&out.DiscoveryStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMeshPeerStatus.
func (in *ServiceMeshPeerStatus) DeepCopy() *ServiceMeshPeerStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceMeshPeerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceName) DeepCopyInto(out *ServiceName) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceName.
func (in *ServiceName) DeepCopy() *ServiceName {
	if in == nil {
		return nil
	}
	out := new(ServiceName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceNameMapping) DeepCopyInto(out *ServiceNameMapping) {
	*out = *in
	out.ServiceName = in.ServiceName
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(ServiceName)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceNameMapping.
func (in *ServiceNameMapping) DeepCopy() *ServiceNameMapping {
	if in == nil {
		return nil
	}
	out := new(ServiceNameMapping)
	in.DeepCopyInto(out)
	return out
}

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	corev1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "maistra.io/api/core/v1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentStatus)(nil), (*DeploymentStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentStatus_To_v1alpha1_DeploymentStatus(a.(*v1.DeploymentStatus), b.(*DeploymentStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DeploymentStatus)(nil), (*v1.DeploymentStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DeploymentStatus_To_v1_DeploymentStatus(a.(*DeploymentStatus), b.(*v1.DeploymentStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ServiceMeshExtension)(nil), (*ServiceMeshExtension)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ServiceMeshExtension_To_v1alpha1_ServiceMeshExtension(a.(*v1.ServiceMeshExtension), b.(*ServiceMeshExtension), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceMeshExtension)(nil), (*v1.ServiceMeshExtension)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ServiceMeshExtension_To_v1_ServiceMeshExtension(a.(*ServiceMeshExtension), b.(*v1.ServiceMeshExtension), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ServiceMeshExtensionList)(nil), (*ServiceMeshExtensionList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ServiceMeshExtensionList_To_v1alpha1_ServiceMeshExtensionList(a.(*v1.ServiceMeshExtensionList), b.(*ServiceMeshExtensionList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceMeshExtensionList)(nil), (*v1.ServiceMeshExtensionList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ServiceMeshExtensionList_To_v1_ServiceMeshExtensionList(a.(*ServiceMeshExtensionList), b.(*v1.ServiceMeshExtensionList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ServiceMeshExtensionStatus)(nil), (*ServiceMeshExtensionStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ServiceMeshExtensionStatus_To_v1alpha1_ServiceMeshExtensionStatus(a.(*v1.ServiceMeshExtensionStatus), b.(*ServiceMeshExtensionStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceMeshExtensionStatus)(nil), (*v1.ServiceMeshExtensionStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ServiceMeshExtensionStatus_To_v1_ServiceMeshExtensionStatus(a.(*ServiceMeshExtensionStatus), b.(*v1.ServiceMeshExtensionStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.WorkloadSelector)(nil), (*WorkloadSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_WorkloadSelector_To_v1alpha1_WorkloadSelector(a.(*v1.WorkloadSelector), b.(*WorkloadSelector), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*WorkloadSelector)(nil), (*v1.WorkloadSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_WorkloadSelector_To_v1_WorkloadSelector(a.(*WorkloadSelector), b.(*v1.WorkloadSelector), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.ServiceMeshExtensionSpec)(nil), (*ServiceMeshExtensionSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ServiceMeshExtensionSpec_To_v1alpha1_ServiceMeshExtensionSpec(a.(*v1.ServiceMeshExtensionSpec), b.(*ServiceMeshExtensionSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*ServiceMeshExtensionSpec)(nil), (*v1.ServiceMeshExtensionSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ServiceMeshExtensionSpec_To_v1_ServiceMeshExtensionSpec(a.(*ServiceMeshExtensionSpec), b.(*v1.ServiceMeshExtensionSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_DeploymentStatus_To_v1alpha1_DeploymentStatus(in *v1.DeploymentStatus, out *DeploymentStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.ContainerSHA256 = in.ContainerSHA256
	out.SHA256 = in.SHA256
	out.URL = in.URL
	out.Message = in.Message
	return nil
}

// Convert_v1_DeploymentStatus_To_v1alpha1_DeploymentStatus is an autogenerated conversion function.
func Convert_v1_DeploymentStatus_To_v1alpha1_DeploymentStatus(in *v1.DeploymentStatus, out *DeploymentStatus, s conversion.Scope) error {
	return autoConvert_v1_DeploymentStatus_To_v1alpha1_DeploymentStatus(in, out, s)
}

func autoConvert_v1alpha1_DeploymentStatus_To_v1_DeploymentStatus(in *DeploymentStatus, out *v1.DeploymentStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.ContainerSHA256 = in.ContainerSHA256
	out.SHA256 = in.SHA256
	out.URL = in.URL
	out.Message = in.Message
	return nil
}

// Convert_v1alpha1_DeploymentStatus_To_v1_DeploymentStatus is an autogenerated conversion function.
func Convert_v1alpha1_DeploymentStatus_To_v1_DeploymentStatus(in *DeploymentStatus, out *v1.DeploymentStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_DeploymentStatus_To_v1_DeploymentStatus(in, out, s)
}

func autoConvert_v1_ServiceMeshExtension_To_v1alpha1_ServiceMeshExtension(in *v1.ServiceMeshExtension, out *ServiceMeshExtension, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ServiceMeshExtensionSpec_To_v1alpha1_ServiceMeshExtensionSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ServiceMeshExtensionStatus_To_v1alpha1_ServiceMeshExtensionStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_ServiceMeshExtension_To_v1alpha1_ServiceMeshExtension is an autogenerated conversion function.
func Convert_v1_ServiceMeshExtension_To_v1alpha1_ServiceMeshExtension(in *v1.ServiceMeshExtension, out *ServiceMeshExtension, s conversion.Scope) error {
	return autoConvert_v1_ServiceMeshExtension_To_v1alpha1_ServiceMeshExtension(in, out, s)
}

func autoConvert_v1alpha1_ServiceMeshExtension_To_v1_ServiceMeshExtension(in *ServiceMeshExtension, out *v1.ServiceMeshExtension, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ServiceMeshExtensionSpec_To_v1_ServiceMeshExtensionSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ServiceMeshExtensionStatus_To_v1_ServiceMeshExtensionStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ServiceMeshExtension_To_v1_ServiceMeshExtension is an autogenerated conversion function.
func Convert_v1alpha1_ServiceMeshExtension_To_v1_ServiceMeshExtension(in *ServiceMeshExtension, out *v1.ServiceMeshExtension, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServiceMeshExtension_To_v1_ServiceMeshExtension(in, out, s)
}

func autoConvert_v1_ServiceMeshExtensionList_To_v1alpha1_ServiceMeshExtensionList(in *v1.ServiceMeshExtensionList, out *ServiceMeshExtensionList, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceMeshExtension, len(*in))
		for i := range *in {
			if err := Convert_v1_ServiceMeshExtension_To_v1alpha1_ServiceMeshExtension(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_ServiceMeshExtensionList_To_v1alpha1_ServiceMeshExtensionList is an autogenerated conversion function.
func Convert_v1_ServiceMeshExtensionList_To_v1alpha1_ServiceMeshExtensionList(in *v1.ServiceMeshExtensionList, out *ServiceMeshExtensionList, s conversion.Scope) error {
	return autoConvert_v1_ServiceMeshExtensionList_To_v1alpha1_ServiceMeshExtensionList(in, out, s)
}

func autoConvert_v1alpha1_ServiceMeshExtensionList_To_v1_ServiceMeshExtensionList(in *ServiceMeshExtensionList, out *v1.ServiceMeshExtensionList, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.ServiceMeshExtension, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_ServiceMeshExtension_To_v1_ServiceMeshExtension(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_ServiceMeshExtensionList_To_v1_ServiceMeshExtensionList is an autogenerated conversion function.
func Convert_v1alpha1_ServiceMeshExtensionList_To_v1_ServiceMeshExtensionList(in *ServiceMeshExtensionList, out *v1.ServiceMeshExtensionList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServiceMeshExtensionList_To_v1_ServiceMeshExtensionList(in, out, s)
}

func autoConvert_v1_ServiceMeshExtensionSpec_To_v1alpha1_ServiceMeshExtensionSpec(in *v1.ServiceMeshExtensionSpec, out *ServiceMeshExtensionSpec, s conversion.Scope) error {
	out.Image = in.Image
	out.ImagePullPolicy = corev1.PullPolicy(in.ImagePullPolicy)
	out.ImagePullSecrets = *(*[]corev1.LocalObjectReference)(unsafe.Pointer(&in.ImagePullSecrets))
	if err := Convert_v1_WorkloadSelector_To_v1alpha1_WorkloadSelector(&in.WorkloadSelector, &out.WorkloadSelector, s); err != nil {
		return err
	}
	out.Phase = (*FilterPhase)(unsafe.Pointer(in.Phase))
	out.Priority = (*int)(unsafe.Pointer(in.Priority))
	// WARNING: in.Config requires manual conversion: inconvertible types (maistra.io/api/core/v1.ServiceMeshExtensionConfig vs string)
	return nil
}

func autoConvert_v1alpha1_ServiceMeshExtensionSpec_To_v1_ServiceMeshExtensionSpec(in *ServiceMeshExtensionSpec, out *v1.ServiceMeshExtensionSpec, s conversion.Scope) error {
	out.Image = in.Image
	out.ImagePullPolicy = corev1.PullPolicy(in.ImagePullPolicy)
	out.ImagePullSecrets = *(*[]corev1.LocalObjectReference)(unsafe.Pointer(&in.ImagePullSecrets))
	if err := Convert_v1alpha1_WorkloadSelector_To_v1_WorkloadSelector(&in.WorkloadSelector, &out.WorkloadSelector, s); err != nil {
		return err
	}
	out.Phase = (*v1.FilterPhase)(unsafe.Pointer(in.Phase))
	out.Priority = (*int)(unsafe.Pointer(in.Priority))
	// WARNING: in.Config requires manual conversion: inconvertible types (string vs maistra.io/api/core/v1.ServiceMeshExtensionConfig)
	return nil
}

func autoConvert_v1_ServiceMeshExtensionStatus_To_v1alpha1_ServiceMeshExtensionStatus(in *v1.ServiceMeshExtensionStatus, out *ServiceMeshExtensionStatus, s conversion.Scope) error {
	out.Phase = FilterPhase(in.Phase)
	out.Priority = in.Priority
	out.ObservedGeneration = in.ObservedGeneration
	if err := Convert_v1_DeploymentStatus_To_v1alpha1_DeploymentStatus(&in.Deployment, &out.Deployment, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_ServiceMeshExtensionStatus_To_v1alpha1_ServiceMeshExtensionStatus is an autogenerated conversion function.
func Convert_v1_ServiceMeshExtensionStatus_To_v1alpha1_ServiceMeshExtensionStatus(in *v1.ServiceMeshExtensionStatus, out *ServiceMeshExtensionStatus, s conversion.Scope) error {
	return autoConvert_v1_ServiceMeshExtensionStatus_To_v1alpha1_ServiceMeshExtensionStatus(in, out, s)
}

func autoConvert_v1alpha1_ServiceMeshExtensionStatus_To_v1_ServiceMeshExtensionStatus(in *ServiceMeshExtensionStatus, out *v1.ServiceMeshExtensionStatus, s conversion.Scope) error {
	out.Phase = v1.FilterPhase(in.Phase)
	out.Priority = in.Priority
	out.ObservedGeneration = in.ObservedGeneration
	if err := Convert_v1alpha1_DeploymentStatus_To_v1_DeploymentStatus(&in.Deployment, &out.Deployment, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ServiceMeshExtensionStatus_To_v1_ServiceMeshExtensionStatus is an autogenerated conversion function.
func Convert_v1alpha1_ServiceMeshExtensionStatus_To_v1_ServiceMeshExtensionStatus(in *ServiceMeshExtensionStatus, out *v1.ServiceMeshExtensionStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServiceMeshExtensionStatus_To_v1_ServiceMeshExtensionStatus(in, out, s)
}

func autoConvert_v1_WorkloadSelector_To_v1alpha1_WorkloadSelector(in *v1.WorkloadSelector, out *WorkloadSelector, s conversion.Scope) error {
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	return nil
}

// Convert_v1_WorkloadSelector_To_v1alpha1_WorkloadSelector is an autogenerated conversion function.
func Convert_v1_WorkloadSelector_To_v1alpha1_WorkloadSelector(in *v1.WorkloadSelector, out *WorkloadSelector, s conversion.Scope) error {
	return autoConvert_v1_WorkloadSelector_To_v1alpha1_WorkloadSelector(in, out, s)
}

func autoConvert_v1alpha1_WorkloadSelector_To_v1_WorkloadSelector(in *WorkloadSelector, out *v1.WorkloadSelector, s conversion.Scope) error {
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	return nil
}

// Convert_v1alpha1_WorkloadSelector_To_v1_WorkloadSelector is an autogenerated conversion function.
func Convert_v1alpha1_WorkloadSelector_To_v1_WorkloadSelector(in *WorkloadSelector, out *v1.WorkloadSelector, s conversion.Scope) error {
	return autoConvert_v1alpha1_WorkloadSelector_To_v1_WorkloadSelector(in, out, s)
}

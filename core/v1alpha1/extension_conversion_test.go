// Copyright Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	corev1 "k8s.io/api/core/v1"
	v1 "maistra.io/api/core/v1"
)

func TestConversion(t *testing.T) {
	var phasev1 v1.FilterPhase = v1.FilterPhasePreAuthN
	var phasev1alpha1 FilterPhase = FilterPhasePreAuthN
	var one int = 1

	cases := []struct {
		name        string
		smev1       v1.ServiceMeshExtension
		smev1alpha1 ServiceMeshExtension
	}{
		{
			name:        "Both Empty",
			smev1:       v1.ServiceMeshExtension{},
			smev1alpha1: ServiceMeshExtension{},
		},
		{
			name: "Empty Config",
			smev1: v1.ServiceMeshExtension{
				Spec: v1.ServiceMeshExtensionSpec{
					Image:    "quay.io/x/y:latest",
					Phase:    &phasev1,
					Priority: &one,
				},
			},
			smev1alpha1: ServiceMeshExtension{
				Spec: ServiceMeshExtensionSpec{
					Image:    "quay.io/x/y:latest",
					Phase:    &phasev1alpha1,
					Priority: &one,
					Config:   "",
				},
			},
		},
		{
			name: "Complete",
			smev1: v1.ServiceMeshExtension{
				Spec: v1.ServiceMeshExtensionSpec{
					Image:            "quay.io/x/y:latest",
					ImagePullPolicy:  corev1.PullNever,
					ImagePullSecrets: []corev1.LocalObjectReference{{Name: "secret1"}, {Name: "secret2"}},
					WorkloadSelector: v1.WorkloadSelector{Labels: map[string]string{"abc": "def"}},
					Phase:            &phasev1,
					Priority:         &one,
					Config: v1.ServiceMeshExtensionConfig{
						Data: map[string]interface{}{
							"string": "some_string",
							"number": float64(10),
							"struct": map[string]interface{}{
								"key": "value",
							},
						},
					},
				},
			},
			smev1alpha1: ServiceMeshExtension{
				Spec: ServiceMeshExtensionSpec{
					Image:            "quay.io/x/y:latest",
					ImagePullPolicy:  corev1.PullNever,
					ImagePullSecrets: []corev1.LocalObjectReference{{Name: "secret1"}, {Name: "secret2"}},
					WorkloadSelector: WorkloadSelector{Labels: map[string]string{"abc": "def"}},
					Phase:            &phasev1alpha1,
					Priority:         &one,
					Config:           "{\"number\":10,\"string\":\"some_string\",\"struct\":{\"key\":\"value\"}}",
				},
			},
		},
		{
			name: "Invalid Config",
			smev1: v1.ServiceMeshExtension{
				Spec: v1.ServiceMeshExtensionSpec{
					Config: v1.ServiceMeshExtensionConfig{
						Data: map[string]interface{}{
							RawV1Alpha1Config: "plain string value",
						},
					},
				},
			},
			smev1alpha1: ServiceMeshExtension{
				Spec: ServiceMeshExtensionSpec{
					Config: "plain string value",
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// v1alpha1 -> v1 -> v1alpha1
			resultV1 := alphav1ToV1(t, c.smev1alpha1, c.smev1)
			v1ToAlphav1(t, resultV1, c.smev1alpha1)

			// v1 -> v1alpha1 -> v1
			resultV1Alpha1 := v1ToAlphav1(t, c.smev1, c.smev1alpha1)
			alphav1ToV1(t, resultV1Alpha1, c.smev1)
		})
	}
}

func v1ToAlphav1(t *testing.T, smev1 v1.ServiceMeshExtension, smev1alpha1 ServiceMeshExtension) (dst ServiceMeshExtension) {
	smev1alpha1.DeepCopyInto(&dst)
	if err := dst.ConvertFrom(&smev1); err != nil {
		t.Fatalf("error converting from v1 to v1alpha1: %v", err)
	}
	if diff := cmp.Diff(smev1alpha1, dst); diff != "" {
		t.Fatalf("conversion diff from v1 to v1alpha1:\n%s", diff)
	}

	return
}

func alphav1ToV1(t *testing.T, smev1alpha1 ServiceMeshExtension, smev1 v1.ServiceMeshExtension) (dst v1.ServiceMeshExtension) {
	if err := smev1alpha1.ConvertTo(&dst); err != nil {
		t.Fatalf("error converting from v1alpha1 to v1: %v", err)
	}
	if diff := cmp.Diff(smev1, dst); diff != "" {
		t.Fatalf("conversion diff from v1alpha1 to v1:\n%s", diff)
	}

	return
}

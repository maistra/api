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

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"maistra.io/api/client/versioned"
	corev1 "maistra.io/api/core/v1"
)

func createClientSet() *versioned.Clientset {
	kubeconfig := os.Getenv("KUBECONFIG")

	if len(kubeconfig) == 0 {
		log.Fatalf("Environment variable KUBECONFIG needs to be set")
	}

	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Failed to create k8s rest client: %s", err)
	}

	return versioned.NewForConfigOrDie(restConfig)
}

func listExtensions(cs *versioned.Clientset) {
	fmt.Printf("Listing Extensions in all namespaces:\n")
	list, err := cs.CoreV1().ServiceMeshExtensions(metav1.NamespaceAll).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error listing: %v\n", err)
		return
	}

	if len(list.Items) == 0 {
		fmt.Printf("No extensions found.\n")
		return
	}

	for _, e := range list.Items {
		fmt.Printf("- Found extension %s/%s (ready: %v)\n", e.Namespace, e.Name, e.Status.Deployment.Ready)
		fmt.Printf("\tConfig for this extension:\n")

		// e is of type v1.ServiceMeshExtension
		// the copy below is not needed, it just ilustrates an use of the api
		var copy *corev1.ServiceMeshExtension = e.DeepCopy()
		config, err := copy.Spec.Config.MarshalJSON()
		if err != nil {
			fmt.Printf("\terror getting extension config: %v\n", err)
		} else {
			fmt.Printf("\t%s\n", config)
		}

	}
}

func listControlPlanes(cs *versioned.Clientset) {
	fmt.Printf("Listing Control planes in all namespaces:\n")
	list, err := cs.CoreV2().ServiceMeshControlPlanes(metav1.NamespaceAll).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error listing: %v\n", err)
		return
	}

	if len(list.Items) == 0 {
		fmt.Printf("No control planes found.")
		return
	}

	for _, cp := range list.Items {
		fmt.Printf("- Found control plane %s/%s: (version: %s)\n", cp.Namespace, cp.Name, cp.Status.ChartVersion)

		memberroll, err := cs.CoreV1().ServiceMeshMemberRolls(cp.Namespace).Get(context.TODO(), "default", metav1.GetOptions{})
		if err != nil {
			fmt.Printf("could not get the `default' SMMR in the %s namespace: %v\n", cp.Namespace, err)
			continue
		}

		fmt.Printf("\tMembers of this control plane:\n")
		for _, member := range memberroll.Spec.Members {
			fmt.Printf("\t\t- %s\n", member)
		}

	}
}

func main() {
	cs := createClientSet()

	listExtensions(cs)
	fmt.Printf("\n")
	listControlPlanes(cs)
}

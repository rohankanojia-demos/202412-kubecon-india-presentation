package golang

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"testing"
)

func TestListPods(t *testing.T) {
	// Use kubeconfig from home directory by default
	kubeconfig := ""
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = home + "/.kube/config"
	}

	// If running in a pod, you may want to use in-cluster config instead of kubeconfig file.
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	// Create a clientset to interact with the cluster
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes clientset: %v", err)
	}

	// Retrieve the list of Pods from the default namespace
	podsClient := clientset.CoreV1().Pods("default") // specify the namespace
	pods, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing pods: %v", err)
	}

	// Print the names of all Pods
	fmt.Println("Listing Pods in 'default' namespace:")
	for _, pod := range pods.Items {
		fmt.Printf("- %s\n", pod.Name)
	}
}

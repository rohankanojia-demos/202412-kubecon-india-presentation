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
	rawConfig, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		log.Fatalf("Error loading kubeconfig file: %v", err)
	}
	currentContext := rawConfig.CurrentContext
	currentContextConfig, ok := rawConfig.Contexts[currentContext]
	if !ok {
		log.Fatalf("Current context %q not fonud", currentContext)
	}
	namespace := currentContextConfig.Namespace
	if namespace == "" {
		namespace = "default"
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

	pods := clientset.CoreV1().Pods(namespace)
	podsList, err := pods.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing pods: %v", err)
	}

	// Print the names of all Pods
	fmt.Println("Listing Pods in current namespace:")
	for _, pod := range podsList.Items {
		fmt.Printf("- %s\n", pod.Name)
	}
}

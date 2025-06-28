package golang

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"os"
	"sigs.k8s.io/yaml"
	"testing"
)

func TestLoadDeploymentYaml(t *testing.T) {
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

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	deployment, err := readDeploymentYAML("../../artifacts/deployment.yaml")
	if err != nil {
		log.Fatal(err)
	}

	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	_, err = deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("Deployment created successfully!")
}

func readDeploymentYAML(filename string) (*appsv1.Deployment, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML file: %w", err)
	}

	var deployment appsv1.Deployment
	err = yaml.Unmarshal(data, &deployment)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return &deployment, nil
}

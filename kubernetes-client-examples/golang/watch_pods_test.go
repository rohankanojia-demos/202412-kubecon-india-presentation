package golang

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"testing"
	"time"
)

func TestWatchPods(t *testing.T) {
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

	w, err := clientset.CoreV1().Pods(namespace).Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event := <-w.ResultChan():
			switch event.Type {
			case watch.Added:
				fmt.Printf("Pod %s/%s was added\n", event.Object.(*corev1.Pod).Namespace, event.Object.(*corev1.Pod).Name)
			case watch.Deleted:
				fmt.Printf("Pod %s/%s was deleted\n", event.Object.(*corev1.Pod).Namespace, event.Object.(*corev1.Pod).Name)
			case watch.Modified:
				fmt.Printf("Pod %s/%s was modified\n", event.Object.(*corev1.Pod).Namespace, event.Object.(*corev1.Pod).Name)
			default:
				fmt.Printf("Unknown event type: %v\n", event.Type)
			}
		case <-time.After(10 * time.Second):
			// Add a timeout to prevent blocking indefinitely
			fmt.Println("Timeout reached, exiting...")
			return
		}
	}
	defer w.Stop()
}

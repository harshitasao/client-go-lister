package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	// NOTE: choose proper kuberntes version and client-go version which are compatible

	//getting the kubeconfig file
	kubeconfig := flag.String("kubeconfig", "~/.kube/config", "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {

	}
	// creating clientset which is going to interact with different resources present in the different api versions
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {

	}
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	for _, pod := range pods.Items {
		fmt.Printf("%s", pod.Name)
	}
	fmt.Println(pods)

}

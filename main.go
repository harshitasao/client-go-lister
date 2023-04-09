package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	// NOTE: choose proper kuberntes version and client-go version which are compatible

	//getting the kubeconfig file
	kubeconfig := flag.String("kubeconfig", "/home/harshitasao/.kube/config", "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("error %s building config from flag\n", err.Error())
		// using the default service account tokens via the inclusterconfig to make our app able top talk to k8s-api-server
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error %s in getting inclusterconfig\n", err.Error())
		}
	}
	// creating clientset which is going to interact with different resources present in the different api versions
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("error %s building clientset from config\n", err.Error())
	}
	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error %s while listing all the pods from default namespace\n", err.Error())
	}
	fmt.Println("Pods from default namespace")
	for _, pod := range pods.Items {
		fmt.Printf("%s\n", pod.Name)
	}
	deployments, err := clientset.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error %s while listing all the deployments from default namespace\n", err.Error())
	}
	fmt.Println("deployments from default namespace")
	for _, deployment := range deployments.Items {
		fmt.Printf("%s\n", deployment.Name)
	}
}

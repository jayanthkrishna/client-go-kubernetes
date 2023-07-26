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

	kubeconfig := flag.String("kubeconfig", "/home/eangrki/.kube/config", "location to your config file")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		fmt.Print(err)
	}

	clientSet, err := kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Print(err)
	}

	ctx := context.Background()
	pods, err := clientSet.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})

	if err != nil {
		fmt.Print(err)
	}

	fmt.Print("Pods from default namespace")

	for _, i := range pods.Items {
		fmt.Printf("%s\n", i.Name)
	}

	fmt.Print("Deployments from default namespace")

	deployments, err := clientSet.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})

	if err != nil {
		fmt.Print(err)
	}

	for _, i := range deployments.Items {
		fmt.Printf("%s\n", i.Name)
	}
}

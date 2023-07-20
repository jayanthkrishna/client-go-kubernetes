package main

import (
	"flag"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	kubeconfig := flag.String("kubeconfig", "~/.kube/config", "location to your config file")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	clientSet, err := kubernetes.NewForConfig(config)

	if err != nil {
		print(err)
	}

	fmt.Print(config)

}

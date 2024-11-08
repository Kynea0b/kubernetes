package main

import (
	"context"
	"flag"
	"fmt"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getkubeconfig() *string {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	return kubeconfig
}

func newClientset(kubeconfig *string) *kubernetes.Clientset {

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)

	return clientset
}

func main() {
	var kubeconfig = getkubeconfig()

	// create the clientset
	clientset := newClientset(kubeconfig)

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range pods.Items {
		fmt.Printf("pod name: %s.\n", pod.Name)

	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	if err != nil {
		log.Fatal(err)
	}

}

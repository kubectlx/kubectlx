package ctx

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetKubeClientWithContext() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", GetKubeconfig())
	if err != nil {
		return nil, err
	}
	if clientset, err := kubernetes.NewForConfig(config); err != nil {
		return nil, err
	} else if _, err := clientset.ServerVersion(); err != nil {
		return nil, err
	} else {
		return clientset, nil
	}
}

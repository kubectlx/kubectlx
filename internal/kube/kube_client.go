package kube

import (
	"github.com/cxweilai/kubectlx/internal/ctx"
	"golang.org/x/net/context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetKubeClientWithContext() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", ctx.GetKubeconfig())
	if err != nil {
		return nil, err
	}
	if clientSet, err := kubernetes.NewForConfig(config); err != nil {
		return nil, err
	} else if _, err := clientSet.ServerVersion(); err != nil {
		return nil, err
	} else {
		return clientSet, nil
	}
}

func GetAllNamespace(ctx context.Context) []string {
	var namespaceNames []string
	clientSet, err := GetKubeClientWithContext()
	if err != nil {
		return namespaceNames
	}
	nsClient := clientSet.CoreV1().Namespaces()
	namespaces, err := nsClient.List(ctx, v1.ListOptions{})
	if err != nil {
		return namespaceNames
	}
	for _, ns := range namespaces.Items {
		namespaceNames = append(namespaceNames, ns.Name)
	}
	return namespaceNames
}

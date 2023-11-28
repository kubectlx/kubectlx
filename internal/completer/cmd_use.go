package completer

import (
	"context"
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/io"
	"os"
	"strings"
)

func NewUseCommand() *command.Command {
	return &command.Command{
		Name:        "use",
		Description: "使用该命令切换集群或Namespace",
		Commands: []*command.Command{
			{
				Name:           "cluster",
				Description:    "切换集群",
				DynamicCommand: listKubeconfig,
			},
			{
				Name:           "namespace",
				Description:    "切换Namespace",
				DynamicCommand: listNamespace,
			},
		},
	}
}

func listKubeconfig() []string {
	dir := ctx.GetKubeHome()
	return rangeListKubeconfig(dir)
}

func rangeListKubeconfig(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return []string{}
	}
	var list []string
	for _, entrie := range entries {
		if entrie.IsDir() {
			if entrie.Name() == "cache" {
				continue
			}
			list = append(list, rangeListKubeconfig(dir+"/"+entrie.Name())...)
			continue
		}
		f, err := os.Open(dir + "/" + entrie.Name())
		if err != nil {
			continue
		}
		head, _ := io.ReadAtMost(f, 10)
		if !strings.Contains(string(head), "apiVersion") {
			continue
		}
		list = append(list, dir+"/"+entrie.Name())
	}
	return list
}

func listNamespace() []string {
	namespaceNames := []string{
		"-A",
	}
	clientSet, err := ctx.GetKubeClientWithContext()
	if err != nil {
		return namespaceNames
	}
	nsClient := clientSet.CoreV1().Namespaces()
	namespaces, err := nsClient.List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return namespaceNames
	}
	for _, ns := range namespaces.Items {
		namespaceNames = append(namespaceNames, ns.Name)
	}
	return namespaceNames
}

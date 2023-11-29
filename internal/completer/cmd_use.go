package completer

import (
	"context"
	"fmt"
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"github.com/cxweilai/kubectlx/internal/kube"
	"k8s.io/utils/io"
	"os"
	"strings"
)

func NewUseCommand() *command.Command {
	return &command.Command{
		Name:        "use",
		Description: "使用该命令切换集群或Namespace",
		Func: func(cmd *command.ExecCmd) {
			cmd.Command.Help()
		},
		Commands: []*command.Command{
			{
				Name:        "cluster",
				Description: "切换集群",
				DynamicCommand: &command.DynamicCommand{
					Func:        listKubeconfig,
					Description: "kubeconfig文件路径",
				},
				Func: func(cmd *command.ExecCmd) {
					if cmd.Child == nil {
						cmd.Command.Help()
						return
					}
					// DynamicCommand
					ctx.SetKubeconfig(cmd.Child.Name)
					fmt.Println("success")
				},
			},
			{
				Name:        "namespace",
				Description: "切换Namespace",
				DynamicCommand: &command.DynamicCommand{
					Func: func() []string {
						namespaces := []string{
							"-A",
						}
						return append(namespaces, kube.GetAllNamespace(context.TODO())...)
					},
					Description: "namespace的名称",
				},
				Func: func(cmd *command.ExecCmd) {
					if cmd.Child == nil && len(cmd.Params) == 0 {
						cmd.Command.Help()
						return
					}
					_, err := kube.GetKubeClientWithContext()
					if err != nil {
						fmt.Println("命令执行失败，连接集群失败。")
						return
					}
					// DynamicCommand
					if cmd.Child != nil {
						ctx.SetKubeconfig(cmd.Child.Name)
					} else if _, ok := cmd.GetParam("A"); ok {
						ctx.SetNamespace("*")
					}
					fmt.Println("success")
				},
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

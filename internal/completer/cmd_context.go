package completer

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"k8s.io/utils/io"
	"os"
	"strings"
)

func NewContextCommand() *command.Command {
	return &command.Command{
		Name:        "context",
		Description: "显示当前会话的配置",
		Run: func(cmd *command.ExecCmd) {
			ctx.ShowCtxInfo()
		},
	}
}

func NewUseCommand() *command.Command {
	return &command.Command{
		Name:        "use",
		Description: "使用该命令切换集群或Namespace",
		Commands: []*command.Command{
			{
				Name:        "cluster",
				Description: "切换集群",
				DynamicParam: &command.DynamicParam{
					Func:        listKubeconfig,
					Flag:        "KUBECONFIG_FILE_PATH",
					Description: "kubeconfig文件路径",
				},
				Run: func(cmd *command.ExecCmd) {
					if cmd.Child == nil {
						cmd.Command.Help()
						return
					}
					// DynamicCommand
					if err := ctx.SetKubeconfig(cmd.Child.Name); err != nil {
						fmt.Println("the kubeconfig is not available: " + err.Error())
					} else {
						fmt.Println("success")
					}
				},
			},
			{
				Name:        "namespace",
				Description: "切换Namespace",
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []string {
						namespaces := []string{
							"-A",
						}
						return append(namespaces, ctx.GetAllNamespace()...)
					},
					Flag:        "NAMESPACE_NAME",
					Description: "namespace的名称",
				},
				Run: func(cmd *command.ExecCmd) {
					if cmd.Child == nil && len(cmd.Params) == 0 {
						cmd.Command.Help()
						return
					}
					// DynamicCommand
					if cmd.Child != nil {
						ctx.SetNamespace(cmd.Child.Name)
					} else if _, ok := cmd.GetParam("-A"); ok {
						ctx.SetNamespace("*")
					}
					fmt.Println("success")
				},
			},
		},
	}
}

func listKubeconfig(input string) []string {
	dir := ctx.GetHome() + "/.kube"
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

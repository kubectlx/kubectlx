package completer

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"github.com/cxweilai/kubectlx/internal/kubecli"
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
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					if err := ctx.SetKubeconfig(cmd.Param); err != nil {
						fmt.Println("the kubeconfig is not available: " + err.Error())
					} else {
						fmt.Println("success")
					}
				}),
			},
			{
				Name:        "namespace",
				Description: "切换Namespace",
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetNamespaces()
					},
					Flag:        "NAMESPACE_NAME",
					Description: "namespace的名称",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					if ctx.SetNamespace(cmd.Param) {
						fmt.Println("success")
					} else {
						fmt.Println("error: the namespace " + cmd.Param + " does not exist")
					}
				}),
			},
		},
	}
}

func listKubeconfig(input string) []*command.Param {
	dir := ctx.GetHome() + "/.kube"
	return rangeListKubeconfig(dir)
}

func rangeListKubeconfig(dir string) []*command.Param {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return []*command.Param{}
	}
	var result []*command.Param
	for _, entrie := range entries {
		if entrie.IsDir() {
			if entrie.Name() == "cache" {
				continue
			}
			result = append(result, rangeListKubeconfig(dir+"/"+entrie.Name())...)
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
		result = append(result, &command.Param{
			Name:        dir + "/" + entrie.Name(),
			Description: "",
		})
	}
	return result
}

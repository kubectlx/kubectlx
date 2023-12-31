package completer

import (
	"fmt"
	"github.com/kubectlx/kubectlx/internal/command"
	"github.com/kubectlx/kubectlx/internal/ctx"
	"github.com/kubectlx/kubectlx/internal/kubecli"
)

func NewKubeLogsCommand() *command.Command {
	return &command.Command{
		Name:        "logs",
		Description: "查看容器日记",
		Options: []*command.Option{
			{
				Name:        "-c",
				Description: "查看某个容器的日志，-c后面跟容器的名称",
			},
			{
				Name:        "-f",
				Description: "以日记流实时输出日记",
			},
		},
		DynamicParam: &command.DynamicParam{
			Func: func(input string) []*command.Param {
				return kubecli.GetPods(ctx.GetNamespace(), input, LIMIT_SUGGEST)
			},
			Flag:        "POD_NAME",
			Description: "Pod名",
		},
		Run: WarpHelp(func(cmd *command.ExecCmd) {
			if _, ok := cmd.GetOption("-f"); ok {
				fmt.Println("error: log stream output is not supported yet")
			} else {
				execKubectl(cmd.Input)
			}
		}),
	}
}

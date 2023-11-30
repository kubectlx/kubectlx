package completer

import (
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
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
				return ctx.GetPods(input, 5)
			},
			Flag:        "POD_NAME",
			Description: "Pod名",
		},
		Run: WarpHelp(func(cmd *command.ExecCmd) {
			execKubectl(cmd.Input)
		}),
	}
}

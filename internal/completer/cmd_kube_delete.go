package completer

import (
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"github.com/cxweilai/kubectlx/internal/kubecli"
	"strings"
)

func NewKubeDeleteCommand() *command.Command {
	options := []*command.Option{
		{
			Name:        "--force=true",
			Description: "强制删除",
		},
	}
	return &command.Command{
		Name:        "delete",
		Description: "删除资源",
		DynamicCommands: func() []*command.Command {
			var cmds []*command.Command
			for _, rcmd := range kubecli.GetK8sResourceCommand("删除%s资源") {
				finalCrd := *rcmd
				cmds = append(cmds, &command.Command{
					Name:        finalCrd.Name,
					Description: finalCrd.Description,
					DynamicParam: &command.DynamicParam{
						Func: func(input string) []*command.Param {
							return kubecli.GetK8sResource(finalCrd.Extended["group"], finalCrd.Extended["version"], finalCrd.Name,
								ctx.GetNamespace(), input, LIMIT_SUGGEST)
						},
						Flag:        strings.ToUpper(finalCrd.Name) + "_NAME",
						Description: "删除" + finalCrd.Name,
					},
					Run: WarpHelp(func(cmd *command.ExecCmd) {
						execKubectl(cmd.Input)
					}),
					Options: options,
				})
			}
			for _, crd := range kubecli.GetCrdCommand("删除%s资源") {
				finalCrd := *crd
				cmds = append(cmds, &command.Command{
					Name:        finalCrd.Name,
					Description: finalCrd.Description,
					DynamicParam: &command.DynamicParam{
						Func: func(input string) []*command.Param {
							return kubecli.GetCrdResource(finalCrd.Extended["group"], finalCrd.Extended["version"], finalCrd.Name,
								ctx.GetNamespace(), input, LIMIT_SUGGEST)
						},
						Flag:        strings.ToUpper(finalCrd.Name) + "_NAME",
						Description: "删除" + finalCrd.Name,
					},
					Options: options,
					Run: WarpHelp(func(cmd *command.ExecCmd) {
						execKubectl(cmd.Input)
					}),
				})
			}
			return cmds
		},
	}
}

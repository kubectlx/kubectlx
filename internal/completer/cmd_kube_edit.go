package completer

import (
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"github.com/cxweilai/kubectlx/internal/kubecli"
	"strings"
)

func NewKubeEditCommand() *command.Command {
	return &command.Command{
		Name:        "edit",
		Description: "编辑资源",
		DynamicCommands: func() []*command.Command {
			var cmds []*command.Command
			for _, rcmd := range kubecli.GetK8sResourceCommand("编辑%s资源") {
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
						Description: "编辑" + finalCrd.Name,
					},
					Run: WarpHelp(func(cmd *command.ExecCmd) {
						execKubectl(cmd.Input)
					}),
				})
			}
			for _, crd := range kubecli.GetCrdCommand("编辑%s资源") {
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
						Description: "编辑" + finalCrd.Name,
					},
					Run: WarpHelp(func(cmd *command.ExecCmd) {
						execKubectl(cmd.Input)
					}),
				})
			}
			return cmds
		},
	}
}

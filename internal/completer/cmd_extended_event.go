package completer

import (
	"fmt"
	"github.com/kubectlx/kubectlx/internal/command"
	"github.com/kubectlx/kubectlx/internal/ctx"
	"github.com/kubectlx/kubectlx/internal/kubecli"
	"strings"
)

func NewExtendedEventCommand() *command.Command {
	return &command.Command{
		Name:        "event",
		Description: "查询资源的事件",
		DynamicCommands: func() []*command.Command {
			var cmds []*command.Command
			for _, rcmd := range kubecli.GetK8sResourceCommand("查询某个%s资源的事件") {
				if rcmd.Name == "configmaps" || rcmd.Name == "secrets" {
					continue
				}
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
						Description: "查询某个" + finalCrd.Name + "资源的事件",
					},
					Run: WarpHelp(func(cmd *command.ExecCmd) {
						execEventCommand(cmd)
					}),
				})
			}
			for _, crd := range kubecli.GetCrdCommand("查询某个%s资源的事件") {
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
						Description: "查询某个" + finalCrd.Name + "资源的事件",
					},
					Run: WarpHelp(func(cmd *command.ExecCmd) {
						execEventCommand(cmd)
					}),
				})
			}
			return cmds
		},
	}
}

func execEventCommand(cmd *command.ExecCmd) {
	realCmd := fmt.Sprintf("get events --field-selector involvedObject.name=%s", cmd.Param)
	execKubectl(realCmd)
}

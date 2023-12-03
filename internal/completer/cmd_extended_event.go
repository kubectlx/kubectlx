package completer

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"github.com/cxweilai/kubectlx/internal/kubecli"
	"strings"
)

func NewExtendedEventCommand() *command.Command {
	return &command.Command{
		Name:        "event",
		Description: "查询资源的事件",
		Commands: []*command.Command{
			{
				Name:        "pods",
				Description: "查询某个Pod资源的事件",
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetPods(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "POD_NAME",
					Description: "Pod名称",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execEventCommand(cmd)
				}),
			},
			{
				Name:        "deployments",
				Description: "查询某个Deployment资源的事件",
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetDeployments(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "DEPLOYMENT_NAME",
					Description: "Deployment名称",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execEventCommand(cmd)
				}),
			},
			{
				Name:        "statefulsets",
				Description: "查询某个StatefulSet资源的事件",
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetStatefulSets(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "STATEFUL_SET_NAME",
					Description: "StatefulSet名称",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execEventCommand(cmd)
				}),
			},
		},
		DynamicCommands: func() []*command.Command {
			var cmds []*command.Command
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

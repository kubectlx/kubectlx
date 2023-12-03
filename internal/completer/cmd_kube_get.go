package completer

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"github.com/cxweilai/kubectlx/internal/kubecli"
	"strings"
)

func NewKubeGetCommand() *command.Command {
	options := []*command.Option{
		{
			Name:        "-A",
			Description: "扫描所有Namespace",
		},
		{
			Name:        "-o",
			Description: "输出格式：json|yaml|wide，例如：-o yaml",
		},
	}
	return &command.Command{
		Name:        "get",
		Description: "资源查询",
		Commands: []*command.Command{
			{
				Name:        "crds",
				Description: typeDescription("查询%s资源", "CRD"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetCrdDefinitions(input, LIMIT_SUGGEST)
					},
					Flag:        "CRD_NAME",
					Description: "回车查看所有CRD，或指定CRD名称可查看指定CRD",
				},
				Run: WarpIgnoreNilParamHelp(func(cmd *command.ExecCmd) {
					execGetOrListResource(cmd, nil)
				}),
				Options: options,
			},
			{
				Name:        "events",
				Description: typeDescription("查询%s资源", "Event"),
				Options: []*command.Option{
					{
						Name:        "--field-selector",
						Description: "用于获取指定资源对象的事件，例如：--field-selector involvedObject.name=<resource name>",
					},
				},
				Run: WarpIgnoreNilParamHelp(func(cmd *command.ExecCmd) {
					execGetOrListResource(cmd, nil)
				}),
			},
		},
		DynamicCommands: func() []*command.Command {
			var cmds []*command.Command
			for _, rcmd := range kubecli.GetK8sResourceCommand("查询%s资源") {
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
						Description: "回车查看所有" + finalCrd.Name + "，或指定" + finalCrd.Name + "名称可查看指定" + finalCrd.Name,
					},
					Run: WarpHelp(func(cmd *command.ExecCmd) {
						execGetOrListResource(cmd, &finalCrd)
					}),
					Options: options,
				})
			}
			for _, crd := range kubecli.GetCrdCommand("查询%s资源") {
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
						Description: "回车查看所有" + finalCrd.Name + "，或指定" + finalCrd.Name + "名称可查看指定" + finalCrd.Name,
					},
					Options: options,
					Run: WarpIgnoreNilParamHelp(func(cmd *command.ExecCmd) {
						execGetOrListResource(cmd, &finalCrd)
					}),
				})
			}
			return cmds
		},
	}
}

func execGetOrListResource(cmd *command.ExecCmd, crd *command.Param) {
	// list操作支持-A获取所有namespace下的资源
	if _, ok := cmd.GetOption("-A"); ok && cmd.Param != "" {
		fmt.Println("error: -A only supports get the resource list")
		return
	}
	execKubectl(cmd.Input)
}

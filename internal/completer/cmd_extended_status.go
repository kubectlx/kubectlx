package completer

import (
	"fmt"
	"github.com/kubectlx/kubectlx/internal/command"
	"github.com/kubectlx/kubectlx/internal/ctx"
	"github.com/kubectlx/kubectlx/internal/kubecli"
	"sigs.k8s.io/yaml"
	"strings"
)

func NewExtendedStatusCommand() *command.Command {
	options := []*command.Option{
		{
			Name:        "--fc",
			Description: "过滤conditions：<字段名>=<字段值>，支持模糊匹配，例如：message=%xx%，即过滤获取message字段包含字符串'xx'的condition。%xx-以xx结尾，xx%-以xx开头, %xx%-包含xx",
		},
		{
			Name:        "--c",
			Description: "结果只显示conditions，其它忽略",
		},
		{
			Name:        "--nc",
			Description: "结果不显示conditions，只显示其它字段",
		},
	}
	return &command.Command{
		Name:        "status",
		Description: "查询资源状态详细信息",
		DynamicCommands: func() []*command.Command {
			var cmds []*command.Command
			for _, rcmd := range kubecli.GetK8sResourceCommand("查询%s资源状态") {
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
						Description: "查询" + finalCrd.Name + "状态，资源名称支持前缀模糊搜索。",
					},
					Run: WarpHelp(func(cmd *command.ExecCmd) {
						execStatusCommand(cmd, &finalCrd)
					}),
					Options: options,
				})
			}
			for _, crd := range kubecli.GetCrdCommand("查询%s资源状态") {
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
						Description: "查询" + finalCrd.Name + "状态，资源名称支持前缀模糊搜索。",
					},
					Run: WarpHelp(func(cmd *command.ExecCmd) {
						execStatusCommand(cmd, &finalCrd)
					}),
					Options: options,
				})
			}
			return cmds
		},
	}
}

func execStatusCommand(cmd *command.ExecCmd, crd *command.Param) {
	var result []*kubecli.ResourceStatus
	if !kubecli.IsK8sResource(crd.Extended["group"], crd.Name) {
		result = kubecli.SearchCrdResourceStatus(ctx.GetNamespace(), crd.Extended["group"], crd.Extended["version"], crd.Name, cmd.Param)
	} else {
		result = kubecli.SearchK8sResourceStatus(ctx.GetNamespace(), crd.Extended["group"], crd.Extended["version"], cmd.Command.Name, cmd.Param)
	}
	// 只取conditions字段
	if _, ok := cmd.GetOption("--c"); ok {
		for _, r := range result {
			s := map[string]interface{}{}
			for k, v := range r.Status.(map[string]interface{}) {
				if k != "conditions" {
					continue
				}
				s[k] = v
			}
			r.Status = s
		}
	}
	// 忽略conditions字段
	if _, ok := cmd.GetOption("--nc"); ok {
		for _, r := range result {
			if _, found := r.Status.(map[string]interface{})["conditions"]; found {
				delete(r.Status.(map[string]interface{}), "conditions")
			}
		}
	}
	if fc, ok := cmd.GetOption("--fc"); ok && strings.Contains(fc, "=") {
		fcArr := strings.Split(fc, "=")
		fieldName := fcArr[0]
		fieldValue := fcArr[1]
		if fieldValue != "" {
			for _, r := range result {
				if conditions, found := r.Status.(map[string]interface{})["conditions"]; found {
					newConditions := make([]interface{}, 0)
					if conditionArr, convOk := conditions.([]interface{}); convOk {
						for _, c := range conditionArr {
							condition := c.(map[string]interface{})
							if fv, fieldFound := condition[fieldName]; fieldFound {
								if value, sok := fv.(string); sok {
									if strings.HasPrefix(fieldValue, "%") && strings.HasSuffix(fieldValue, "%") {
										if strings.Contains(value, fieldValue[1:len(fieldValue)-1]) {
											newConditions = append(newConditions, condition)
										}
									} else if strings.HasPrefix(fieldValue, "%") {
										if strings.HasSuffix(value, fieldValue[1:]) {
											newConditions = append(newConditions, condition)
										}
									} else if strings.HasSuffix(fieldValue, "%") {
										if strings.HasPrefix(value, fieldValue[:len(fieldValue)-1]) {
											newConditions = append(newConditions, condition)
										}
									} else if value == fieldValue {
										newConditions = append(newConditions, condition)
									}
								}
							}
						}
					}
					r.Status.(map[string]interface{})["conditions"] = newConditions
				}
			}
		}
	}
	uiResult := map[string]interface{}{}
	uiResult["Type"] = cmd.Command.Name
	uiResult["Results"] = result
	statusJson, _ := yaml.Marshal(uiResult)
	fmt.Println(string(statusJson))
}

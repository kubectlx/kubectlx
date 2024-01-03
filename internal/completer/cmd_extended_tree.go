package completer

import (
	"fmt"
	"github.com/kubectlx/kubectlx/internal/command"
	"github.com/kubectlx/kubectlx/internal/ctx"
	"github.com/kubectlx/kubectlx/internal/kubecli"
)

func NewExtendedTreeCommand() *command.Command {
	return &command.Command{
		Name:        "tree",
		Description: "查询资源以及关联的'子资源'",
		Commands: []*command.Command{
			{
				Name:        "deployments",
				Description: "查询Deployment资源树",
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetK8sResource("apps", "v1", "deployments", ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "DEPLOYMENT_NAME",
					Description: "指定Deployment资源的名称",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					tree, err := kubecli.GetDeploymentTree(ctx.GetNamespace(), cmd.Param)
					if err != nil {
						fmt.Println("error: " + err.Error())
						return
					}
					showTree(tree, 0)
				}),
			},
			{
				Name:        "services",
				Description: "查询Service资源树",
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetK8sResource("", "v1", "services", ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "SERVICE_NAME",
					Description: "指定Service资源的名称",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					tree, err := kubecli.GetServiceTree(ctx.GetNamespace(), cmd.Param)
					if err != nil {
						fmt.Println("error: " + err.Error())
						return
					}
					showTree(tree, 0)
				}),
			},
			{
				Name:        "statefulsets",
				Description: "查询StatefulSet资源树",
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetK8sResource("apps", "v1", "statefulsets", ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "STATEFUL_SET_NAME",
					Description: "指定StatefulSet资源的名称",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					tree, err := kubecli.GetStatefulSetTree(ctx.GetNamespace(), cmd.Param)
					if err != nil {
						fmt.Println("error: " + err.Error())
						return
					}
					showTree(tree, 0)
				}),
			},
			{
				Name:        "daemonsets",
				Description: "查询DaemonSet资源树",
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetK8sResource("apps", "v1", "daemonsets", ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "DAEMON_SET_NAME",
					Description: "指定DaemonSet资源的名称",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					tree, err := kubecli.GetDaemonSetTree(ctx.GetNamespace(), cmd.Param)
					if err != nil {
						fmt.Println("error: " + err.Error())
						return
					}
					showTree(tree, 0)
				}),
			},
			{
				Name:        "jobs",
				Description: "查询Job资源树",
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetK8sResource("batch", "v1", "jobs", ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "DEPLOYMENT_NAME",
					Description: "指定Job资源的名称",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					tree, err := kubecli.GetJobTree(ctx.GetNamespace(), cmd.Param)
					if err != nil {
						fmt.Println("error: " + err.Error())
						return
					}
					showTree(tree, 0)
				}),
			},
		},
	}
}

func showTree(tree *kubecli.ResourceTree, level int) {
	spacePrefix := ""
	for i := 0; i < level; i++ {
		spacePrefix += "  |----"
	}
	fmt.Println(fmt.Sprintf("%s%s\t%s\t%s", spacePrefix, tree.Type, tree.Name, tree.Status))
	if len(tree.Child) == 0 {
		return
	}
	for _, child := range tree.Child {
		showTree(child, level+1)
	}
}

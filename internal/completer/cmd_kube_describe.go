package completer

import (
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"github.com/cxweilai/kubectlx/internal/kubecli"
	"strings"
)

func NewKubeDescribeCommand() *command.Command {
	return &command.Command{
		Name:        "describe",
		Description: "查看资源详情",
		Commands: []*command.Command{
			{
				Name:        "crds",
				Description: typeDescription("查看%s资源详情", "CRD"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetCrdDefinitions(input, 5)
					},
					Flag:        "CRD_NAME",
					Description: "查看CRD资源详情",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "namespaces",
				Description: typeDescription("查看%s资源详情", "Namespace"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetNamespaces()
					},
					Flag:        "NAMESPACE_NAME",
					Description: "查看Namespace资源详情",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "nodes",
				Description: typeDescription("查看%s资源详情", "Node"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetNodes(input, 5)
					},
					Flag:        "NODE_NAME",
					Description: "查看Node资源详情",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "pods",
				Description: typeDescription("查看%s资源详情", "Pod"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetPods(ctx.GetNamespace(), input, 5)
					},
					Flag:        "POD_NAME",
					Description: "查看Pod资源详情",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "services",
				Description: typeDescription("查看%s资源详情", "Service"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetServices(ctx.GetNamespace(), input, 5)
					},
					Flag:        "SERVICE_NAME",
					Description: "查看Service资源详情",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "deployments",
				Description: typeDescription("查看%s资源详情", "Deployment"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetDeployments(ctx.GetNamespace(), input, 5)
					},
					Flag:        "DEPLOYMENT_NAME",
					Description: "查看Deployment资源详情",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "daemonsets",
				Description: typeDescription("查看%s资源详情", "DaemonSet"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetDaemonSets(ctx.GetNamespace(), input, 5)
					},
					Flag:        "DAEMON_SET_NAME",
					Description: "查看DaemonSet资源详情",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "replicasets",
				Description: typeDescription("查看%s资源详情", "ReplicaSet"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetReplicaSets(ctx.GetNamespace(), input, 5)
					},
					Flag:        "REPLICA_SET_NAME",
					Description: "查看ReplicaSet资源详情",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "statefulsets",
				Description: typeDescription("查看%s资源详情", "StatefulSet"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetStatefulSets(ctx.GetNamespace(), input, 5)
					},
					Flag:        "STATEFUL_SET_NAME",
					Description: "查看StatefulSet资源详情",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "jobs",
				Description: typeDescription("查看%s资源详情", "Job"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetJobs(ctx.GetNamespace(), input, 5)
					},
					Flag:        "JOB_NAME",
					Description: "查看Job资源详情",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "cronjobs",
				Description: typeDescription("查看%s资源详情", "CronJob"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetCronJobs(ctx.GetNamespace(), input, 5)
					},
					Flag:        "CRON_JOB_NAME",
					Description: "查看CronJob资源详情",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "configmaps",
				Description: typeDescription("查看%s资源详情", "ConfigMap"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetConfigMaps(ctx.GetNamespace(), input, 5)
					},
					Flag:        "CONFIG_MAP_NAME",
					Description: "查看ConfigMap资源详情",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "secrets",
				Description: typeDescription("查看%s资源详情", "Secret"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetSecrets(ctx.GetNamespace(), input, 5)
					},
					Flag:        "SECRET_NAME",
					Description: "查看Secret资源详情",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
		},
		DynamicCommands: func() []*command.Command {
			var cmds []*command.Command
			for _, crd := range kubecli.GetCrdCommand("获取%s资源详情") {
				finalCrd := *crd
				cmds = append(cmds, &command.Command{
					Name:        finalCrd.Name,
					Description: finalCrd.Description,
					DynamicParam: &command.DynamicParam{
						Func: func(input string) []*command.Param {
							return kubecli.GetCrdResource(finalCrd.Extended["group"], finalCrd.Extended["version"], finalCrd.Name,
								ctx.GetNamespace(), input, 5)
						},
						Flag:        strings.ToUpper(finalCrd.Name) + "_NAME",
						Description: "查看" + finalCrd.Name + "资源详情",
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

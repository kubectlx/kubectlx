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
		Commands: []*command.Command{
			{
				Name:        "crds",
				Description: typeDescription("删除%s资源", "CRD"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetCrdDefinitions(input, LIMIT_SUGGEST)
					},
					Flag:        "CRD_NAME",
					Description: "删除CRD",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
				Options: options,
			},
			{
				Name:        "namespaces",
				Description: typeDescription("删除%s", "Namespace"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetNamespaces()
					},
					Flag:        "NAMESPACE_NAME",
					Description: "删除Namespace",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
				Options: options,
			},
			{
				Name:        "pods",
				Description: typeDescription("删除%s资源", "Pod"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetPods(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "POD_NAME",
					Description: "删除Pod",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
				Options: options,
			},
			{
				Name:        "services",
				Description: typeDescription("删除%s资源", "Service"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetServices(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "SERVICE_NAME",
					Description: "删除Service",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
				Options: options,
			},
			{
				Name:        "deployments",
				Description: typeDescription("删除%s资源", "Deployment"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetDeployments(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "DEPLOYMENT_NAME",
					Description: "删除Deployment",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
				Options: options,
			},
			{
				Name:        "daemonsets",
				Description: typeDescription("删除%s资源", "DaemonSet"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetDaemonSets(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "DAEMON_SET_NAME",
					Description: "删除DaemonSet",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
				Options: options,
			},
			{
				Name:        "replicasets",
				Description: typeDescription("删除%s资源", "ReplicaSet"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetReplicaSets(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "REPLICA_SET_NAME",
					Description: "删除ReplicaSet",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
				Options: options,
			},
			{
				Name:        "statefulsets",
				Description: typeDescription("删除%s资源", "StatefulSet"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetStatefulSets(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "STATEFUL_SET_NAME",
					Description: "删除StatefulSet",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
				Options: options,
			},
			{
				Name:        "jobs",
				Description: typeDescription("删除%s资源", "Job"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetJobs(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "JOB_NAME",
					Description: "删除Job",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
				Options: options,
			},
			{
				Name:        "cronjobs",
				Description: typeDescription("删除%s资源", "CronJob"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetCronJobs(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "CRON_JOB_NAME",
					Description: "删除CronJob",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
				Options: options,
			},
			{
				Name:        "configmaps",
				Description: typeDescription("删除%s资源", "ConfigMap"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetConfigMaps(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "CONFIG_MAP_NAME",
					Description: "删除ConfigMap",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
				Options: options,
			},
			{
				Name:        "secrets",
				Description: typeDescription("删除%s资源", "Secret"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetSecrets(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "SECRET_NAME",
					Description: "删除Secret",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
				Options: options,
			},
		},
		DynamicCommands: func() []*command.Command {
			var cmds []*command.Command
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

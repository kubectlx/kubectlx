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
		Commands: []*command.Command{
			{
				Name:        "pods",
				Description: typeDescription("编辑%s资源", "Pod"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetPods(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "POD_NAME",
					Description: "编辑Pod",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "services",
				Description: typeDescription("编辑%s资源", "Service"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetServices(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "SERVICE_NAME",
					Description: "编辑Service",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "deployments",
				Description: typeDescription("编辑%s资源", "Deployment"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetDeployments(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "DEPLOYMENT_NAME",
					Description: "编辑Deployment",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "daemonsets",
				Description: typeDescription("编辑%s资源", "DaemonSet"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetDaemonSets(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "DAEMON_SET_NAME",
					Description: "编辑DaemonSet",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "replicasets",
				Description: typeDescription("编辑%s资源", "ReplicaSet"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetReplicaSets(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "REPLICA_SET_NAME",
					Description: "编辑ReplicaSet",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "statefulsets",
				Description: typeDescription("编辑%s资源", "StatefulSet"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetStatefulSets(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "STATEFUL_SET_NAME",
					Description: "编辑StatefulSet",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "jobs",
				Description: typeDescription("编辑%s资源", "Job"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetJobs(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "JOB_NAME",
					Description: "编辑job",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "cronjobs",
				Description: typeDescription("编辑%s资源", "CronJob"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetCronJobs(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "CRON_JOB_NAME",
					Description: "编辑CronJob",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "configmaps",
				Description: typeDescription("编辑%s资源", "ConfigMap"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetConfigMaps(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "CONFIG_MAP_NAME",
					Description: "编辑ConfigMap",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
			{
				Name:        "secrets",
				Description: typeDescription("编辑%s资源", "Secret"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetSecrets(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "SECRET_NAME",
					Description: "编辑Secret",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execKubectl(cmd.Input)
				}),
			},
		},
		DynamicCommands: func() []*command.Command {
			var cmds []*command.Command
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

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
				Name:        "namespaces",
				Description: typeDescription("查询%s资源", "Namespace"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetNamespaces()
					},
					Flag:        "NAMESPACE_NAME",
					Description: "回车查看所有Namespace，或指定Namespace名称可查看指定Namespace",
				},
				Run: WarpIgnoreNilParamHelp(func(cmd *command.ExecCmd) {
					execGetOrListResource(cmd, nil)
				}),
				Options: options,
			},
			{
				Name:        "nodes",
				Description: typeDescription("查询%s资源", "Node"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetNodes(input, LIMIT_SUGGEST)
					},
					Flag:        "NODE_NAME",
					Description: "回车查看所有Node，或指定Node名称可查看指定Node",
				},
				Run: WarpIgnoreNilParamHelp(func(cmd *command.ExecCmd) {
					execGetOrListResource(cmd, nil)
				}),
				Options: options,
			},
			{
				Name:        "pods",
				Description: typeDescription("查询%s资源", "Pod"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetPods(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "POD_NAME",
					Description: "回车查看所有Pod，或指定pod名称可查看指定Pod",
				},
				Run: WarpIgnoreNilParamHelp(func(cmd *command.ExecCmd) {
					execGetOrListResource(cmd, nil)
				}),
				Options: options,
			},
			{
				Name:        "services",
				Description: typeDescription("查询%s资源", "Service"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetServices(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "SERVICE_NAME",
					Description: "回车查看所有Service，或指定Service名称可查看指定Service",
				},
				Run: WarpIgnoreNilParamHelp(func(cmd *command.ExecCmd) {
					execGetOrListResource(cmd, nil)
				}),
				Options: options,
			},
			{
				Name:        "deployments",
				Description: typeDescription("查询%s资源", "Deployment"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetDeployments(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "DEPLOYMENT_NAME",
					Description: "回车查看所有Deployment，或指定Deployment名称可查看指定Deployment",
				},
				Run: WarpIgnoreNilParamHelp(func(cmd *command.ExecCmd) {
					execGetOrListResource(cmd, nil)
				}),
				Options: options,
			},
			{
				Name:        "daemonsets",
				Description: typeDescription("查询%s资源", "DaemonSet"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetDaemonSets(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "DAEMON_SET_NAME",
					Description: "回车查看所有DaemonSet，或指定DaemonSet名称可查看指定DaemonSet",
				},
				Run: WarpIgnoreNilParamHelp(func(cmd *command.ExecCmd) {
					execGetOrListResource(cmd, nil)
				}),
				Options: options,
			},
			{
				Name:        "replicasets",
				Description: typeDescription("查询%s资源", "ReplicaSet"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetReplicaSets(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "REPLICA_SET_NAME",
					Description: "回车查看所有ReplicaSet，或指定ReplicaSet名称可查看指定ReplicaSet",
				},
				Run: WarpIgnoreNilParamHelp(func(cmd *command.ExecCmd) {
					execGetOrListResource(cmd, nil)
				}),
				Options: options,
			},
			{
				Name:        "statefulsets",
				Description: typeDescription("查询%s资源", "StatefulSet"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetStatefulSets(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "STATEFUL_SET_NAME",
					Description: "回车查看所有StatefulSet，或指定StatefulSet名称可查看指定StatefulSet",
				},
				Run: WarpIgnoreNilParamHelp(func(cmd *command.ExecCmd) {
					execGetOrListResource(cmd, nil)
				}),
				Options: options,
			},
			{
				Name:        "jobs",
				Description: typeDescription("查询%s资源", "Job"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetJobs(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "JOB_NAME",
					Description: "回车查看所有job，或指定job名称可查看指定job",
				},
				Run: WarpIgnoreNilParamHelp(func(cmd *command.ExecCmd) {
					execGetOrListResource(cmd, nil)
				}),
				Options: options,
			},
			{
				Name:        "cronjobs",
				Description: typeDescription("查询%s资源", "CronJob"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetCronJobs(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "CRON_JOB_NAME",
					Description: "回车查看所有CronJob，或指定CronJob名称可查看指定CronJob",
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
			{
				Name:        "configmaps",
				Description: typeDescription("查询%s资源", "ConfigMap"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetConfigMaps(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "CONFIG_MAP_NAME",
					Description: "回车查看所有ConfigMap，或指定ConfigMap名称可查看指定ConfigMap",
				},
				Run: WarpIgnoreNilParamHelp(func(cmd *command.ExecCmd) {
					execGetOrListResource(cmd, nil)
				}),
				Options: options,
			},
			{
				Name:        "secrets",
				Description: typeDescription("查询%s资源", "Secret"),
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetSecrets(ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "SECRET_NAME",
					Description: "回车查看所有Secret，或指定Secret名称可查看指定Secret",
				},
				Run: WarpIgnoreNilParamHelp(func(cmd *command.ExecCmd) {
					execGetOrListResource(cmd, nil)
				}),
				Options: options,
			},
		},
		DynamicCommands: func() []*command.Command {
			var cmds []*command.Command
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

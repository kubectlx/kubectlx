package completer

import "github.com/kubectlx/kubectlx/internal/command"

func NewKubeApplyCommand() *command.Command {
	return &command.Command{
		Name:        "apply",
		Description: "通过yaml文件来创建或更新Kubernetes资源对象",
		Options: []*command.Option{
			{
				Name:        "-f",
				Description: "<file.yaml>|<directory> -f参数可以指定包含资源配置的YAML文件的路径，或者包含多个资源配置文件的目录路径，apply将会根据文件中的配置来创建或更新相应的资源",
			},
			{
				Name:        "--dry-run=client",
				Description: "只对资源配置进行客户端验证，不会实际执行操作。可用于验证资源配置的正确性。",
			},
		},
		Run: WarpIgnoreNilParamHelp(func(cmd *command.ExecCmd) {
			if file, ok := cmd.GetOption("-f"); !ok || file == "" {
				cmd.Command.Help()
				return
			}
			execKubectl(cmd.Input)
		}),
	}
}

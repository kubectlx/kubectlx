package completer

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"github.com/cxweilai/kubectlx/internal/kubecli"
	"os"
	"os/exec"
)

func NewKubeExecCommand() *command.Command {
	return &command.Command{
		Name:        "exec",
		Description: "执行容器中的命令",
		Options: []*command.Option{
			{
				Name:        "-c",
				Description: "指定容器名，如果省略，则会选择Pod中的第一个容器。",
			},
			{
				Name:        "-it",
				Description: "-it是-i和-t的组合，一般都是组合使用。-t：分配一个交互式终端来执行命令。-i：保持开启标准输入，并且可以接收来自用户的输入。",
			},
			{
				Name:        "--",
				Description: "执行的命令，例如：-- /bin/bash 或 -- /bin/sh",
			},
		},
		DynamicParam: &command.DynamicParam{
			Func: func(input string) []*command.Param {
				return kubecli.GetPods(ctx.GetNamespace(), input, 5)
			},
			Flag:        "POD_NAME",
			Description: "pod的名称",
		},
		Run: WarpHelp(func(cmd *command.ExecCmd) {
			flags := fmt.Sprintf("--kubeconfig=%s -n %s", ctx.GetKubeconfig(), ctx.GetNamespace())
			options := ""
			shellCmd := ""
			for k, v := range cmd.Options {
				if k == "--" {
					shellCmd = fmt.Sprintf(" %s %s", k, v)
					continue
				}
				if v != "" {
					options += fmt.Sprintf(" %s %s", k, v)
				} else {
					options += fmt.Sprintf(" %s", k)
				}
			}
			realCmd := fmt.Sprintf("kubectl %s %s %s%s%s", cmd.Name, cmd.Param, flags, options, shellCmd)
			c := exec.Command("/bin/sh", "-c", realCmd)
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			_ = c.Run()
		}),
	}
}

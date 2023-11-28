package executor

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/completer"
	"github.com/cxweilai/kubectlx/internal/ctx"
)

type ctxExecutor struct {
	subCmdRun map[string]func(cmd *command.ExecCmd)
}

func newCtxExecutor() *ctxExecutor {
	return &ctxExecutor{subCmdRun: map[string]func(cmd *command.ExecCmd){
		"cluster": func(cmd *command.ExecCmd) {
			if cmd.Child == nil {
				fmt.Println("未指定Kubeconfig文件路径")
				return
			}
			// DynamicCommand
			ctx.SetKubeconfig(cmd.Child.Name)
			fmt.Println("success")
		},
		"namespace": func(cmd *command.ExecCmd) {
			_, err := ctx.GetKubeClientWithContext()
			if err != nil {
				fmt.Println("命令执行失败，连接集群失败。")
				return
			}
			// DynamicCommand
			if cmd.Child != nil {
				ctx.SetKubeconfig(cmd.Child.Name)
			} else if _, ok := cmd.GetParam("A"); ok {
				ctx.SetNamespace("*")
			}
			fmt.Println("success")
		},
	}}
}

func (c *ctxExecutor) help() {
	cmd := completer.NewUseCommand()
	cmd.Help()
}

func (c *ctxExecutor) Support(cmd *command.ExecCmd) bool {
	if cmd.Name == "use" {
		return true
	}
	return false
}

func (c *ctxExecutor) DoExecutor(cmd *command.ExecCmd, input string) {
	subCmd := cmd.Child
	if subCmd == nil {
		c.help()
		return
	}
	if runer, ok := c.subCmdRun[subCmd.Name]; ok {
		runer(subCmd)
	}
}

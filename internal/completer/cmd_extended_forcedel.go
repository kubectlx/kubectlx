package completer

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"github.com/cxweilai/kubectlx/internal/kubecli"
	"strings"
	"sync"
	"time"
)

func NewExtendedForceDelCommand() *command.Command {
	return &command.Command{
		Name:        "forcedel",
		Description: "强制删除资源（先执行删除，如果对象没删掉，接着移除Finalizers）",
		Commands:    nil,
		DynamicCommands: func() []*command.Command {
			var cmds []*command.Command
			for _, rcmd := range kubecli.GetK8sResourceCommand("强制删除%s资源") {
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
						Description: "强制删除" + finalCrd.Name + "资源",
					},
					Run: WarpHelp(func(cmd *command.ExecCmd) {
						execForceDeleteCommand(cmd)
					}),
				})
			}
			for _, crd := range kubecli.GetCrdCommand("强制删除%s资源") {
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
						Description: "强制删除" + finalCrd.Name + "资源",
					},
					Run: WarpHelp(func(cmd *command.ExecCmd) {
						execForceDeleteCommand(cmd)
					}),
				})
			}
			return cmds
		},
	}
}

func execForceDeleteCommand(cmd *command.ExecCmd) {
	wait := sync.WaitGroup{}
	wait.Add(1)
	go func() {
		select {
		case <-time.After(3 * time.Second):
			execKubectlIgnoreOutput(fmt.Sprintf("patch %s %s -p '{\"metadata\":{\"finalizers\":[]}}' --type=merge", cmd.Command.Name, cmd.Param))
			wait.Done()
		}
	}()
	execKubectl(fmt.Sprintf("delete %s %s", cmd.Command.Name, cmd.Param))
	wait.Wait()
}

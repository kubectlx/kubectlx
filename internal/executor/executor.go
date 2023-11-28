package executor

import (
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/completer"
	"strings"
)

type Executor interface {
	Support(cmd *command.ExecCmd) bool
	DoExecutor(cmd *command.ExecCmd, input string)
}

func RegistryExecutor(e Executor) {
	executors = append(executors, e)
}

var executors []Executor

func init() {
	executors = append(executors, &defaultExecutor{}, &systemExecutor{}, newCtxExecutor())
}

func CmdExecutor(cmd string) {
	execCmd := command.ParseCommand(strings.TrimSpace(cmd))
	if execCmd == nil {
		completer.NewCompleter().Help()
		return
	}

	for i := len(executors) - 1; i >= 0; i-- {
		executor := executors[i]
		if executor.Support(execCmd) {
			executor.DoExecutor(execCmd, cmd)
			return
		}
	}
}

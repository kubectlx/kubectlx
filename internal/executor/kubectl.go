package executor

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"os"
	"os/exec"
)

type defaultExecutor struct {
}

func (d defaultExecutor) Support(cmd *command.ExecCmd) bool {
	return true
}

func (d defaultExecutor) DoExecutor(cmd *command.ExecCmd, input string) {
	exec.Command("/bin/sh", "-c", "export KUBECONFIG=", ctx.GetKubeconfig())
	c := exec.Command("/bin/sh", "-c", "kubectl "+input)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		fmt.Printf("executor command fail: %s\n", err.Error())
	}
	return
}

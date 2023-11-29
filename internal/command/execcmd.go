package command

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"os"
	"os/exec"
	"strings"
)

// ExecCmd include DynamicCommand
type ExecCmd struct {
	Name    string
	Child   *ExecCmd
	Parent  *ExecCmd
	Params  map[string]string
	Command *Command
	Input   string
}

func (ec *ExecCmd) Exec() {
	if ec.Command != nil {
		if ec.Command.Run == nil {
			// 使用默认的kubectl执行
			ec.execKubectl()
			return
		}
		ec.Command.Run(ec)
		return
	}
	if ec.Parent != nil {
		// 交给父命令执行
		ec.Parent.Exec()
	}
}

func (ec *ExecCmd) execKubectl() {
	exec.Command("/bin/sh", "-c", "export KUBECONFIG=", ctx.GetKubeconfig())
	c := exec.Command("/bin/sh", "-c", "kubectl "+ec.Input)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		fmt.Printf("executor command fail: %s\n", err.Error())
	}
}

// GetParam 支持忽略'-'前缀
func (ec *ExecCmd) GetParam(name string) (string, bool) {
	if ec.Params == nil {
		return "", false
	}
	if strings.HasPrefix(name, "-") {
		v, ok := ec.Params["-"+name]
		return v, ok
	} else {
		v, ok := ec.Params[name]
		return v, ok
	}
}

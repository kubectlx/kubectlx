package completer

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"os"
	"os/exec"
)

func execKubectl(input string) {
	var flags string
	if ctx.GetNamespace() == "*" {
		flags = fmt.Sprintf("--kubeconfig=%s -A", ctx.GetKubeconfig())
	} else {
		flags = fmt.Sprintf("--kubeconfig=%s -n %s", ctx.GetKubeconfig(), ctx.GetNamespace())
	}
	c := exec.Command("/bin/sh", "-c", "kubectl "+input+" "+flags)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	_ = c.Run()
}

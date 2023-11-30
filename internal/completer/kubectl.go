package completer

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"os"
	"os/exec"
)

func execKubectl(input string) {
	flags := fmt.Sprintf("--kubeconfig=%s -n %s", ctx.GetKubeconfig(), ctx.GetNamespace())
	c := exec.Command("/bin/sh", "-c", "kubectl "+input+" "+flags)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	_ = c.Run()
}

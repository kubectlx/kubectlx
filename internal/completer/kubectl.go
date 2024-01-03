package completer

import (
	"fmt"
	"github.com/kubectlx/kubectlx/internal/ctx"
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

type DevNullWrite struct {
}

func (w *DevNullWrite) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func execKubectlIgnoreOutput(input string) {
	flags := fmt.Sprintf("--kubeconfig=%s -n %s", ctx.GetKubeconfig(), ctx.GetNamespace())
	c := exec.Command("/bin/sh", "-c", "kubectl "+input+" "+flags)
	c.Stdin = os.Stdin
	c.Stdout = &DevNullWrite{}
	c.Stderr = &DevNullWrite{}
	_ = c.Run()
}

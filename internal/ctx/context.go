package ctx

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/option"
	"os"
	"runtime"
)

type xContext struct {
	kubeconfig string
	namespace  string
}

var ctx = xContext{}

func InitWithConfig(cfg *option.Config) {
	if cfg.Kubeconfig != "" {
		ctx.kubeconfig = cfg.Kubeconfig
		fmt.Println("use kubeconfig: " + cfg.Kubeconfig)
	}
}

func SetKubeconfig(kubeconfig string) {
	ctx.kubeconfig = kubeconfig
}

func GetKubeHome() string {
	switch runtime.GOOS {
	case "windows":
		return os.Getenv("USERPROFILE") + "/.kube"
	case "linux", "darwin":
		return os.Getenv("HOME") + "/.kube"
	}
	return ""
}

func GetKubeconfig() string {
	if ctx.kubeconfig == "" {
		if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
			return kubeconfig
		}
		return GetKubeHome() + "/config"
	}
	return ctx.kubeconfig
}

func SetNamespace(namespace string) {
	ctx.namespace = namespace
}

func GetNamespace() string {
	if ctx.namespace == "" {
		return "default"
	}
	return ctx.namespace
}

func ShowCtxInfo() {
	fmt.Println("use cluster: " + GetKubeconfig())
	fmt.Println("use namespace: " + GetNamespace())
}

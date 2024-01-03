package ctx

import (
	"fmt"
	"github.com/kubectlx/kubectlx/internal/kubecli"
	"github.com/kubectlx/kubectlx/internal/option"
	"log"
	"os"
	"runtime"
	"strings"
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
	} else {
		ctx.kubeconfig = getKubeconfig()
		fmt.Println("default use kubeconfig: " + ctx.kubeconfig)
		fmt.Print("you can specify the kubeconfig using the '--kubeconfig' parameter: ")
		fmt.Println("--kubeconfig=~/.kube/config_k3s")
	}
	// 初始化失败退出进程
	if err := kubecli.InitKubeClient(GetKubeconfig()); err != nil {
		log.Fatalln(err)
	}
	if SetNamespace(cfg.Namespace) {
		fmt.Println("use namespace: " + GetNamespace())
	} else {
		fmt.Println("default use namespace: " + GetNamespace())
		fmt.Print("you can specify the namespace using the '--namespace' parameter: ")
		fmt.Println("--namespace=<your namespace>")
	}
}

func SetKubeconfig(kubeconfig string) error {
	if err := kubecli.InitKubeClient(replaceKubeconfigHomePath(kubeconfig)); err != nil {
		return err
	}
	ctx.kubeconfig = kubeconfig
	return nil
}

func GetHome() string {
	switch runtime.GOOS {
	case "windows":
		return os.Getenv("USERPROFILE")
	case "linux", "darwin":
		return os.Getenv("HOME")
	}
	return ""
}

func GetKubeconfig() string {
	kubeconfig := getKubeconfig()
	return replaceKubeconfigHomePath(kubeconfig)
}

func getKubeconfig() string {
	if ctx.kubeconfig == "" {
		if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
			return kubeconfig
		}
		return GetHome() + "/.kube/config"
	}
	return ctx.kubeconfig
}

func replaceKubeconfigHomePath(kubeconfig string) string {
	if strings.HasPrefix(kubeconfig, "~/") {
		return strings.ReplaceAll(kubeconfig, "~", GetHome())
	}
	return kubeconfig
}

func SetNamespace(namespace string) bool {
	found := false
	namespaces := kubecli.GetNamespaces()
	for _, ns := range namespaces {
		if ns.Name == namespace {
			found = true
			break
		}
	}
	if !found {
		return false
	}
	ctx.namespace = namespace
	return true
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

package ctx

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/option"
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
	if err := initKubeClientWithContext(GetKubeconfig()); err != nil {
		log.Fatalln(err)
	}
}

func SetKubeconfig(kubeconfig string) error {
	if err := initKubeClientWithContext(replaceKubeconfigHomePath(kubeconfig)); err != nil {
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

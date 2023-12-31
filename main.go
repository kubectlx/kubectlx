package main

import (
	"flag"
	"github.com/c-bata/go-prompt"
	"github.com/kubectlx/kubectlx/internal/completer"
	"github.com/kubectlx/kubectlx/internal/ctx"
	"github.com/kubectlx/kubectlx/internal/option"
)

func main() {
	cfg := &option.Config{}
	flag.StringVar(&cfg.Kubeconfig, "kubeconfig", "", "The kubeconfig the default use k8s cluster kubeconfig.")
	flag.StringVar(&cfg.Namespace, "namespace", "", "The namespace is used to specify the namespace.")
	flag.Parse()
	ctx.InitWithConfig(cfg)

	c := completer.NewCompleter()
	p := prompt.New(c.Exec, c.Complete,
		prompt.OptionTitle("kubectlx: interactive kubernetes client"),
		prompt.OptionPrefix("kubectlx >>> "),
	)
	p.Run()
}

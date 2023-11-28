package main

import (
	"flag"
	"github.com/c-bata/go-prompt"
	"github.com/cxweilai/kubectlx/internal/completer"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"github.com/cxweilai/kubectlx/internal/executor"
	"github.com/cxweilai/kubectlx/internal/option"
)

func main() {
	cfg := &option.Config{}
	flag.StringVar(&cfg.Kubeconfig, "kubeconfig", "", "The kubeconfig the default use k8s cluster kubeconfig.")
	flag.Parse()
	ctx.InitWithConfig(cfg)

	c := completer.NewCompleter()
	p := prompt.New(executor.CmdExecutor, c.Complete,
		prompt.OptionTitle("kubectlx: interactive kubernetes client"),
		prompt.OptionPrefix("kubectlx >>> "),
	)
	p.Run()
}

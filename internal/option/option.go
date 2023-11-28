package option

type Config struct {
	Kubeconfig string
}

type Option func(cfg *Config)

func WithKubeconfig(kubeconfig string) Option {
	return func(cfg *Config) {
		cfg.Kubeconfig = kubeconfig
	}
}

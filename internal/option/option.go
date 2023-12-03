package option

type Config struct {
	Kubeconfig string
	Namespace  string
}

type Option func(cfg *Config)

func WithKubeconfig(kubeconfig string) Option {
	return func(cfg *Config) {
		cfg.Kubeconfig = kubeconfig
	}
}

func WithNamespace(namespace string) Option {
	return func(cfg *Config) {
		cfg.Namespace = namespace
	}
}

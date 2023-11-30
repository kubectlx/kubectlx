package command

import "errors"

// flags是全局的flag，例如 -n namespace (-A不是flag，是get命令的option)，--kubeconfig=xx等，
// 由于-n和--kubeconfig已经抽离独立命令支持，所以隐藏掉了

type Flag struct {
	Name        string
	Description string
}

func (f *Flag) Check() error {
	if f.Name == "" {
		return errors.New("flag name required")
	}
	if f.Description == "" {
		return errors.New("flag description required")
	}
	return nil
}

var flags []*Flag

func init() {
	flags = []*Flag{
		{
			Name:        "-h",
			Description: "查看当前命令帮助文档",
		},
		{
			Name:        "-help",
			Description: "查看当前命令的帮助文档",
		},
	}
	for _, f := range flags {
		if err := f.Check(); err != nil {
			panic(err)
		}
	}
}

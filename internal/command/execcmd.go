package command

import (
	"strings"
)

// ExecCmd include DynamicCommand
type ExecCmd struct {
	Name    string
	Child   *ExecCmd
	Parent  *ExecCmd
	Params  map[string]string
	Command *Command
	Input   string
}

func (ec *ExecCmd) Exec() bool {
	if ec.Command != nil {
		ec.Command.Run(ec)
		return true
	}
	if ec.Parent != nil {
		// 交给父命令执行
		return ec.Parent.Exec()
	}
	return false
}

// GetParam 支持忽略'-'前缀
func (ec *ExecCmd) GetParam(name string) (string, bool) {
	if ec.Params == nil {
		return "", false
	}
	if strings.HasPrefix(name, "-") {
		v, ok := ec.Params["-"+name]
		return v, ok
	} else {
		v, ok := ec.Params[name]
		return v, ok
	}
}

package command

import (
	"strings"
)

type ExecCmd struct {
	Name    string
	Child   *ExecCmd
	Parent  *ExecCmd
	Param   string
	Options map[string]string // 包含flag
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

// GetOption 支持忽略'-'前缀
func (ec *ExecCmd) GetOption(name string) (string, bool) {
	if ec.Options == nil {
		return "", false
	}
	if !strings.HasPrefix(name, "-") {
		v, ok := ec.Options["-"+name]
		return v, ok
	} else {
		v, ok := ec.Options[name]
		return v, ok
	}
}

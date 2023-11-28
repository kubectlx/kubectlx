package command

import (
	"strings"
)

// ExecCmd include DynamicCommand
type ExecCmd struct {
	Name   string
	Child  *ExecCmd
	Params map[string]string
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

func ParseCommand(cmd string) (command *ExecCmd) {
	args := getCmdArgs(cmd)
	return parseCmdAndArgs(args)
}

func getCmdArgs(cmd string) []string {
	args := strings.Split(cmd, " ")
	var fArgs []string
	for _, arg := range args {
		if arg != "" {
			fArgs = append(fArgs, arg)
		}
	}
	return fArgs
}

// '-'开头的一定是配置参数，但参数名后面不一定是参数值，可能参数不需要值，例如：'-it' 是参数 '-n default' 也是参数
// 不带'-'开头的，前面一个是带'-'开头的，那么一定是前面参数的参数值，否则就是命令，
// 例如：kubectl get pod xxx 其中'xxx'不当作参数，而当作命令，指动态命令（DynamicCommand）
func parseCmdAndArgs(args []string) (root *ExecCmd) {
	var (
		command *ExecCmd
		params  map[string]string
	)
	for idx := 0; idx < len(args); {
		arg := args[idx]
		if strings.HasPrefix(arg, "-") {
			if len(args) <= idx+1 || strings.HasPrefix(args[idx+1], "-") {
				command.Params[args[idx]] = ""
				idx += 1
				continue
			}
			command.Params[args[idx]] = args[idx+1]
			idx += 2
			continue
		}
		if command == nil {
			command = &ExecCmd{Name: arg}
			root = command
		} else {
			command.Child = &ExecCmd{Name: arg}
			command = command.Child
		}
		idx++
	}
	if command == nil {
		return nil
	}
	command.Params = params
	return
}

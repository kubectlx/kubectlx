package command

import (
	"strings"
)

func ParseExecCommand(command []*Command, input string) (execCmd *ExecCmd) {
	args := getExecCmdArgs(input)
	execCmd = parseExecCmdAndArgs(args, input)
	if execCmd == nil {
		return nil
	}
	root := execCmd
	for root.Parent != nil {
		root = root.Parent
	}
	bindCommand(command, root)
	return
}

func getExecCmdArgs(cmd string) []string {
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
func parseExecCmdAndArgs(args []string, input string) *ExecCmd {
	var (
		command *ExecCmd
		parent  *ExecCmd
		params  = map[string]string{}
	)
	for idx := 0; idx < len(args); {
		arg := args[idx]
		if strings.HasPrefix(arg, "-") {
			if len(args) <= idx+1 || strings.HasPrefix(args[idx+1], "-") {
				params[args[idx]] = ""
				idx += 1
				continue
			}
			params[args[idx]] = args[idx+1]
			idx += 2
			continue
		}
		if command == nil {
			command = &ExecCmd{Name: arg, Input: input}
			parent = command
		} else {
			command.Child = &ExecCmd{Name: arg, Input: input}
			command = command.Child
			command.Parent = parent
			parent = command
		}
		idx++
	}
	if command == nil {
		return nil
	}
	command.Params = params
	return command
}

func bindCommand(commands []*Command, execCmd *ExecCmd) {
	if len(commands) == 0 || execCmd == nil {
		return
	}
	for _, c := range commands {
		if c.Name == execCmd.Name {
			execCmd.Command = c
			bindCommand(c.Commands, execCmd.Child)
			return
		}
	}
}

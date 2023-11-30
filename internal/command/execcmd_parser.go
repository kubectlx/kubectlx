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

// '-'开头的一定是option，但option后面不一定是option的值，可能option不需要值，例如：'-it'
// 不带'-'开头的，前面一个是带'-'开头的，那么一定是前面option的值，否则就是命令或者参数
// 例如：kubectl get pod xxx 其中'xxx'是动态参数（DynamicParam）
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

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
	rebindParamAndOptionsToLastSubCommand(execCmd)
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
		options = map[string]string{}
	)
	for idx := 0; idx < len(args); {
		arg := args[idx]
		if strings.HasPrefix(arg, "-") {
			if len(args) <= idx+1 || strings.HasPrefix(args[idx+1], "-") {
				options[args[idx]] = ""
				idx += 1
				continue
			}
			options[args[idx]] = args[idx+1]
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
	command.Options = options
	return command
}

// 将执行命令与声明命令关联
func bindCommand(commands []*Command, execCmd *ExecCmd) {
	if len(commands) == 0 || execCmd == nil {
		return
	}
	for _, c := range commands {
		if c.Name == execCmd.Name {
			execCmd.Command = c
			var childCommands []*Command
			if len(c.Commands) > 0 {
				childCommands = append(childCommands, c.Commands...)
			}
			if c.DynamicCommands != nil {
				dCommands := c.DynamicCommands()
				if len(dCommands) > 0 {
					childCommands = append(childCommands, dCommands...)
				}
			}
			bindCommand(childCommands, execCmd.Child)
			return
		}
	}
}

// 将动态参数和Options重新绑定到最后一个子命令
func rebindParamAndOptionsToLastSubCommand(execCmd *ExecCmd) {
	if execCmd.Command != nil {
		return
	}
	ptr := execCmd
	if ptr.Command == nil && ptr.Parent != nil {
		prePtr := ptr.Parent
		prePtr.Options = mergeMap(prePtr.Options, ptr.Options)
		prePtr.Param = ptr.Name
		prePtr.Child = nil
		ptr = prePtr
	}
}

func mergeMap(source map[string]string, target map[string]string) map[string]string {
	if source == nil {
		return target
	}
	if target == nil {
		return source
	}
	// 用target，允许子命令的option覆盖父命令的option
	for k, v := range target {
		source[k] = v
	}
	return source
}

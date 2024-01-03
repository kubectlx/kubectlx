package completer

import "github.com/kubectlx/kubectlx/internal/command"

func WarpHelp(f func(cmd *command.ExecCmd)) func(cmd *command.ExecCmd) {
	return func(cmd *command.ExecCmd) {
		if doHelpIfHelp(cmd) {
			return
		}
		if cmd.Param == "" {
			cmd.Command.Help()
			return
		}
		f(cmd)
	}
}

func WarpIgnoreNilParamHelp(f func(cmd *command.ExecCmd)) func(cmd *command.ExecCmd) {
	return func(cmd *command.ExecCmd) {
		if doHelpIfHelp(cmd) {
			return
		}
		f(cmd)
	}
}

func doHelpIfHelp(cmd *command.ExecCmd) bool {
	_, help := cmd.GetOption("help")
	if !help {
		_, help = cmd.GetOption("h")
	}
	if help {
		cmd.Command.Help()
		return true
	}
	return false
}

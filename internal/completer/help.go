package completer

import "github.com/cxweilai/kubectlx/internal/command"

func WarpHelp(f func(cmd *command.ExecCmd)) func(cmd *command.ExecCmd) {
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
	if help || cmd.Param == "" {
		cmd.Command.Help()
		return true
	}
	return false
}

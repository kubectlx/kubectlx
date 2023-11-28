package completer

import (
	"github.com/cxweilai/kubectlx/internal/command"
)

var systemCommands = map[string]string{
	"quit":    "退出",
	"exit":    "退出",
	"clear":   "清空屏幕",
	"context": "显示当前会话的配置",
}

func NewSystemCommand() []*command.Command {
	var commands []*command.Command
	for cmd, desc := range systemCommands {
		commands = append(commands, &command.Command{Name: cmd, Description: desc})
	}
	return commands
}

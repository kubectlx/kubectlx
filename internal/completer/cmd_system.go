package completer

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"os"
	"os/exec"
	"runtime"
)

func NewSystemCommand() []*command.Command {
	return []*command.Command{
		{
			Name:        "quit",
			Description: "退出",
			Func: func(cmd *command.ExecCmd) {
				fmt.Println("Bye!")
				os.Exit(0)
			},
		},
		{
			Name:        "exit",
			Description: "退出",
			Func: func(cmd *command.ExecCmd) {
				fmt.Println("Bye!")
				os.Exit(0)
			},
		},
		{
			Name:        "clear",
			Description: "清空屏幕",
			Func: func(cmd *command.ExecCmd) {
				_ = clearScreen()
			},
		},
		{
			Name:        "context",
			Description: "显示当前会话的配置",
			Func: func(cmd *command.ExecCmd) {
				ctx.ShowCtxInfo()
			},
		},
	}
}

func clearScreen() error {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		return cmd.Run()
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		return cmd.Run()
	default:
		return nil
	}
}

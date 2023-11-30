package completer

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/command"
	"os"
	"os/exec"
	"runtime"
)

func NewSystemCommand() []*command.Command {
	return []*command.Command{
		{
			Name:        "quit",
			Description: "退出",
			Run: func(cmd *command.ExecCmd) {
				fmt.Println("Bye!")
				os.Exit(0)
			},
			IgnoreFlags: true,
		},
		{
			Name:        "exit",
			Description: "退出",
			Run: func(cmd *command.ExecCmd) {
				fmt.Println("Bye!")
				os.Exit(0)
			},
			IgnoreFlags: true,
		},
		{
			Name:        "clear",
			Description: "清空屏幕",
			Run: func(cmd *command.ExecCmd) {
				_ = clearScreen()
			},
			IgnoreFlags: true,
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

package completer

import (
	"fmt"
	"github.com/kubectlx/kubectlx/internal/command"
	"os"
	"os/exec"
	"runtime"
)

func NewSystemCommand() []*command.Command {
	return []*command.Command{
		{
			Name:        "exit",
			Description: "退出",
			Run: func(cmd *command.ExecCmd) {
				fmt.Println("Bye!")
				os.Exit(0)
			},
		},
		{
			Name:        "clear",
			Description: "清空屏幕",
			Run: func(cmd *command.ExecCmd) {
				_ = clearScreen()
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

package executor

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"os"
)

type systemExecutor struct {
}

func (s *systemExecutor) Support(cmd *command.ExecCmd) bool {
	switch cmd.Name {
	case "quit", "exit", "clear", "context":
		return true
	}
	return false
}

func (s *systemExecutor) DoExecutor(cmd *command.ExecCmd, input string) {
	switch cmd.Name {
	case "quit", "exit":
		fmt.Println("Bye!")
		os.Exit(0)
		return
	case "clean":
		_ = clearScreen()
		return
	case "context":
		ctx.ShowCtxInfo()
	}
}

func clearScreen() error {
	//switch runtime.GOOS {
	//case "windows":
	//	cmd := exec.Command("cmd", "/c", "cls")
	//	cmd.Stdout = os.Stdout
	//	return cmd.Run()
	//case "linux", "darwin":
	//	cmd := exec.Command("clear")
	//	cmd.Stdout = os.Stdout
	//	return cmd.Run()
	//}
	//return nil
	//fmt.Print("\033[H\033[2J")
	return nil
}

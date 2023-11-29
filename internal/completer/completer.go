package completer

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"os"
	"os/exec"
	"strings"
)

type Completer struct {
	cmd *command.Command
}

func NewCompleter() *Completer {
	commands := NewSystemCommand()
	commands = append(commands, NewUseCommand())
	c := &Completer{
		cmd: &command.Command{
			Commands: commands,
		},
	}
	for _, cmd := range c.cmd.Commands {
		if err := cmd.Check(); err != nil {
			panic(err)
		}
	}
	return c
}

func (c *Completer) Registry(cmd ...*command.Command) {
	for _, tcmd := range cmd {
		if err := tcmd.Check(); err != nil {
			panic(err)
		}
	}
	c.cmd.Commands = append(c.cmd.Commands, cmd...)
}

func (c *Completer) Exec(cmd string) {
	execCmd := command.ParseExecCommand(c.cmd.Commands, strings.TrimSpace(cmd))
	if execCmd == nil {
		c.cmd.Help()
		return
	}
	if execCmd.Exec() {
		return
	}
	// c.cmd.Help()
	execKubectl(cmd)
}

func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	args := c.getArgs(d)
	suggests := c.doGetLastCmd(c.cmd, args)
	return prompt.FilterHasPrefix(suggests, d.GetWordBeforeCursor(), true)
}

func (c *Completer) getArgs(d prompt.Document) []string {
	args := strings.Split(d.TextBeforeCursor(), " ")
	var fArgs []string
	for _, arg := range args {
		if arg != "" {
			fArgs = append(fArgs, arg)
		}
	}
	return fArgs
}

func (c *Completer) doGetLastCmd(cmd *command.Command, args []string) []prompt.Suggest {
	if len(args) == 0 {
		return c.getCommandSuggests(cmd)
	}
	firstArg := args[0]
	for _, subCmd := range cmd.Commands {
		if strings.EqualFold(subCmd.Name, firstArg) {
			if len(args) == 1 {
				return c.getCommandSuggests(subCmd)
			}
			return c.doGetLastCmd(subCmd, args[1:])
		}
	}
	return c.getCommandSuggests(cmd)
}

func (c *Completer) getCommandSuggests(cmd *command.Command) []prompt.Suggest {
	if cmd.Commands != nil {
		var suggests []prompt.Suggest
		for _, subCmd := range cmd.Commands {
			suggests = append(suggests, prompt.Suggest{
				Text:        subCmd.Name,
				Description: subCmd.Description,
			})
		}
		return suggests
	} else if cmd.DynamicCommand != nil {
		var suggests []prompt.Suggest
		for _, dp := range cmd.DynamicCommand.Func() {
			suggests = append(suggests, prompt.Suggest{
				Text: dp,
			})
		}
		return suggests
	} else if len(cmd.Args) != 0 {
		var suggests []prompt.Suggest
		for _, option := range cmd.Args {
			suggests = append(suggests, prompt.Suggest{
				Text:        option.Name,
				Description: option.Description,
			})
		}
		return suggests
	}
	return []prompt.Suggest{}
}

func execKubectl(input string) {
	exec.Command("/bin/sh", "-c", "export KUBECONFIG=", ctx.GetKubeconfig())
	c := exec.Command("/bin/sh", "-c", "kubectl "+input)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		fmt.Printf("executor command fail: %s\n", err.Error())
	}
}

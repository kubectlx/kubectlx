package completer

import (
	"github.com/c-bata/go-prompt"
	"github.com/cxweilai/kubectlx/internal/command"
	"strings"
)

type Completer struct {
	cmd *command.Command
}

func NewCompleter() *Completer {
	commands := NewSystemCommand()
	commands = append(commands, NewUseCommand())
	return &Completer{
		cmd: &command.Command{
			Commands: commands,
		},
	}
}

func (c *Completer) Help() {
	c.cmd.Help()
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
		return c.GetCommandSuggests(cmd)
	}
	firstArg := args[0]
	for _, subCmd := range cmd.Commands {
		if strings.EqualFold(subCmd.Name, firstArg) {
			if len(args) == 1 {
				return c.GetCommandSuggests(subCmd)
			}
			return c.doGetLastCmd(subCmd, args[1:])
		}
	}
	return c.GetCommandSuggests(cmd)
}

func (c *Completer) GetCommandSuggests(cmd *command.Command) []prompt.Suggest {
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
		for _, dp := range cmd.DynamicCommand() {
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

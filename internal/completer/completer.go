package completer

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/cxweilai/kubectlx/internal/command"
	"strings"
)

type Completer struct {
	systemCmd    *command.Command
	contextCmd   *command.Command
	kubeCmd      *command.Command
	extensionCmd *command.Command
	cmd          *command.Command
}

func NewCompleter() *Completer {
	contextCmd := &command.Command{
		Name: "Context Command",
		Commands: []*command.Command{
			NewContextCommand(),
			NewUseCommand(),
		},
	}
	systemCmd := &command.Command{
		Name:     "System Command",
		Commands: NewSystemCommand(),
	}
	kubeCmd := &command.Command{
		Name: "Kube Command",
		Commands: []*command.Command{
			NewKubeLogsCommand(),
		},
	}
	var commands []*command.Command
	commands = append(commands, systemCmd.Commands...)
	commands = append(commands, contextCmd.Commands...)
	commands = append(commands, kubeCmd.Commands...)
	c := &Completer{
		contextCmd: contextCmd,
		systemCmd:  systemCmd,
		kubeCmd:    kubeCmd,
		extensionCmd: &command.Command{
			Name:     "Extension Command",
			Commands: []*command.Command{},
		},
		cmd: &command.Command{
			Commands: commands,
		},
	}
	for _, cmd := range c.cmd.Commands {
		if err := cmd.Check(); err != nil {
			panic(err)
		}
	}
	// 追加help命令到头部
	c.cmd.Commands = append([]*command.Command{
		{
			Name:        "help",
			Description: "帮助命令",
			Run: func(cmd *command.ExecCmd) {
				c.Help()
			},
		},
	}, c.cmd.Commands...)
	return c
}

func (c *Completer) Registry(cmd ...*command.Command) {
	for _, tcmd := range cmd {
		if err := tcmd.Check(); err != nil {
			panic(err)
		}
	}
	c.extensionCmd.Commands = append(c.extensionCmd.Commands, cmd...)
	c.cmd.Commands = append(c.cmd.Commands, cmd...)
}

func (c *Completer) Help() {
	c.systemCmd.Help()
	c.contextCmd.Help()
	c.kubeCmd.Help()
	if len(c.extensionCmd.Commands) > 0 {
		c.extensionCmd.Help()
	}
}

func (c *Completer) Exec(cmd string) {
	execCmd := command.ParseExecCommand(c.cmd.Commands, strings.TrimSpace(cmd))
	if execCmd == nil {
		return
	}
	if execCmd.Exec() {
		return
	}
	fmt.Println(fmt.Sprintf("cmd '%s' not found.", cmd))
	fmt.Println("Use the \"help\" command to view supported commands.")
}

func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	// 正常空字符串分割后数组长度为1，元素为空字符串："" => [""]
	args := strings.Split(d.TextBeforeCursor(), " ")
	suggests := c.doGetLastCmd(c.cmd, args)
	return prompt.FilterHasPrefix(suggests, d.GetWordBeforeCursor(), true)
}

func (c *Completer) doGetLastCmd(cmd *command.Command, args []string) []prompt.Suggest {
	if len(args) == 0 {
		return []prompt.Suggest{}
	}
	firstArg := args[0]
	for _, subCmd := range cmd.Commands {
		if strings.EqualFold(subCmd.Name, firstArg) {
			return c.doGetLastCmd(subCmd, args[1:])
		}
	}
	return c.getCommandSuggests(cmd, args)
}

func (c *Completer) getCommandSuggests(cmd *command.Command, args []string) []prompt.Suggest {
	// 1.子命令
	if cmd.Commands != nil {
		// 不合法校验
		// 例如：'c '，输入c之后再输入空格，由于c不匹配，此时args有两个参数，一个是'c'一个是' '，空格之后就不应该模糊匹配了
		if len(args) > 1 {
			return []prompt.Suggest{}
		}
		// 做模糊匹配，避免忽略前面的输入
		// 例如：'xxx '，由于xxx不存在，不过滤的话，空格之后会匹配到根命令，拿根命令的子命令；
		//      'use xxx '，由于xxx不存在，不过滤的话，空格之后会匹配到父命令'use'，拿‘use’的子命令。
		lastArg := ""
		if len(args) > 0 {
			lastArg = args[0]
		}
		var suggests []prompt.Suggest
		for _, subCmd := range cmd.Commands {
			if !strings.HasPrefix(subCmd.Name, lastArg) {
				continue
			}
			suggests = append(suggests, prompt.Suggest{
				Text:        subCmd.Name,
				Description: subCmd.Description,
			})
		}
		return suggests
	}
	// 2.参数
	// len(args) <= 1 ==> 避免参数输入完后还提示输入参数，例如：'logs xxx '，在xxx之后、空格之后，还提示输入参数
	if cmd.DynamicParam != nil && len(args) <= 1 {
		lastArg := ""
		if len(args) > 0 {
			lastArg = args[len(args)-1]
		}
		var suggests []prompt.Suggest
		for _, p := range cmd.DynamicParam.Func(lastArg) {
			suggests = append(suggests, prompt.Suggest{
				Text:        p.Name,
				Description: p.Description,
			})
		}
		return suggests
	}
	// 3. options
	var suggests []prompt.Suggest
	if len(cmd.Options) > 0 {
		for _, option := range cmd.Options {
			suggests = append(suggests, prompt.Suggest{
				Text:        option.Name,
				Description: option.Description,
			})
		}
	}
	return suggests
}

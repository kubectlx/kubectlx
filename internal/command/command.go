package command

import (
	"errors"
	"fmt"
)

type Param struct {
	Name        string
	Description string
}

func (p *Param) Check() error {
	if p.Description == "" {
		return errors.New("param name required")
	}
	if p.Description == "" {
		return errors.New("param description required")
	}
	return nil
}

type DynamicCommand struct {
	Func        func() []string
	Description string
}

func (dc *DynamicCommand) Check() error {
	if dc.Description == "" {
		return errors.New("dynamic command description required")
	}
	if dc.Func == nil {
		return errors.New("dynamic command fun required")
	}
	return nil
}

type Command struct {
	Name           string
	Description    string
	Commands       []*Command
	Args           []*Param
	DynamicCommand *DynamicCommand
	Func           func(cmd *ExecCmd)
}

func (cl *Command) AddCommand(cmds ...*Command) {
	for _, cmd := range cmds {
		cl.Commands = append(cl.Commands, cmd)
	}
}

func (cl *Command) Help() {
	clearLine() // 清除光标所在位置后的一行的标准输出
	if cl.Name == "" {
		fmt.Println("help:")
	} else {
		fmt.Println(cl.Name + " help:")
	}
	if len(cl.Commands) > 0 {
		for _, subCmd := range cl.Commands {
			fmt.Printf("  %s\t%s\n", subCmd.Name, subCmd.Description)
		}
	} else if len(cl.Args) > 0 {
		for _, param := range cl.Args {
			fmt.Printf("  %s\t%s\n\n", param.Name, param.Description)
		}
	} else if cl.DynamicCommand != nil {
		fmt.Printf("  %s\n", cl.DynamicCommand.Description)
	}
}

func clearLine() {
	// 创建并打印包含光标移动到行首和清除行的控制字符序列
	fmt.Print("\r\033[K")
}

func (cl *Command) Check() error {
	if cl.Name == "" {
		return errors.New("command name required")
	}
	if cl.Description == "" {
		return errors.New("command description required")
	}
	if cl.Func == nil {
		if len(cl.Commands) == 0 {
			return errors.New("command func required")
		}
		cl.Func = func(cmd *ExecCmd) {
			cmd.Command.Help()
		}
	}
	if cl.DynamicCommand != nil {
		if err := cl.DynamicCommand.Check(); err != nil {
			return err
		}
	}
	if len(cl.Args) > 0 {
		for _, arg := range cl.Args {
			if err := arg.Check(); err != nil {
				return err
			}
		}
	}
	if len(cl.Commands) > 0 {
		for _, subCmd := range cl.Commands {
			if err := subCmd.Check(); err != nil {
				return err
			}
		}
	}
	return nil
}

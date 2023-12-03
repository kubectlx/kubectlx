package command

import (
	"errors"
	"fmt"
)

type Option struct {
	Name        string
	Description string
}

func (p *Option) Check() error {
	if p.Name == "" {
		return errors.New("option name required")
	}
	if p.Description == "" {
		return errors.New("option description required")
	}
	return nil
}

type Param struct {
	Name        string
	Description string
	Extended    map[string]string
}

type DynamicParam struct {
	// input 当前最后输入的字符串用于支持模糊查询
	// return map[string]string key为值，value为描述
	Func        func(input string) []*Param
	Flag        string
	Description string
}

func (dc *DynamicParam) Check() error {
	if dc.Flag == "" {
		return errors.New("dynamic param flag required")
	}
	if dc.Description == "" {
		return errors.New("dynamic param description required")
	}
	if dc.Func == nil {
		return errors.New("dynamic param fun required")
	}
	return nil
}

type Command struct {
	Name            string
	Description     string
	Commands        []*Command
	DynamicCommands func() []*Command
	Options         []*Option
	DynamicParam    *DynamicParam
	Run             func(cmd *ExecCmd)
}

func (cl *Command) AddCommand(cmds ...*Command) {
	for _, cmd := range cmds {
		cl.Commands = append(cl.Commands, cmd)
	}
}

func (cl *Command) Help() {
	clearLine() // 清除光标所在位置后的一行的标准输出
	fmt.Println(cl.Name + ":")
	if cl.Commands != nil || cl.DynamicCommands != nil {
		if cl.Commands != nil {
			for _, subCmd := range cl.Commands {
				fmt.Printf("  %s\t%s\n", subCmd.Name, subCmd.Description)
			}
		}
		// 支持动态命令
		if cl.DynamicCommands != nil {
			for _, subCmd := range cl.DynamicCommands() {
				fmt.Printf("  %s\t%s\n", subCmd.Name, subCmd.Description)
			}
		}
	} else {
		if cl.DynamicParam != nil {
			fmt.Printf("  (%s)\t%s\n", cl.DynamicParam.Flag, cl.DynamicParam.Description)
		}
		if len(flags) > 0 {
			fmt.Println("  flags:")
			for _, f := range flags {
				fmt.Printf("  %s\t%s\n", f.Name, f.Description)
			}
		}
		if len(cl.Options) > 0 {
			fmt.Println("  options:")
			for _, option := range cl.Options {
				fmt.Printf("    %s\t%s\n", option.Name, option.Description)
			}
		}
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
	if cl.Run == nil {
		// 叶子命令必须指定Run方法
		if cl.Commands == nil && cl.DynamicCommands == nil {
			return errors.New("command func required")
		}
		// 如果还有子命令，则当前命令的Run默认为help命令
		cl.Run = func(cmd *ExecCmd) {
			cmd.Command.Help()
		}
	}
	if cl.DynamicParam != nil {
		if err := cl.DynamicParam.Check(); err != nil {
			return err
		}
	}
	if len(cl.Options) > 0 {
		for _, arg := range cl.Options {
			if err := arg.Check(); err != nil {
				return err
			}
		}
	}
	if cl.Commands != nil {
		for _, subCmd := range cl.Commands {
			if err := subCmd.Check(); err != nil {
				return err
			}
		}
	}
	return nil
}

package command

import "fmt"

type Param struct {
	Name        string
	Description string
}

type DynamicCommand func() []string

type Command struct {
	Name           string
	Description    string
	Commands       []*Command
	Args           []*Param
	DynamicCommand DynamicCommand
}

func (cl *Command) AddCommand(cmds ...*Command) {
	for _, cmd := range cmds {
		cl.Commands = append(cl.Commands, cmd)
	}
}

func (cl *Command) Help() {
	if cl.Commands == nil {
		return
	}
	ClearLine() // 清除光标所在位置后的一行的标准输出
	cl.help("")
}

func (cl *Command) help(psj string) {
	if cl.Commands == nil {
		return
	}
	if cl.Name == "" {
		fmt.Println("help:")
	} else {
		fmt.Println(psj + cl.Name + " help:")
	}
	for _, subCmd := range cl.Commands {
		fmt.Printf(psj+"  %s\t%s\n", subCmd.Name, subCmd.Description)
		if subCmd.Commands != nil {
			subCmd.help(psj + "  ")
		}
	}
}

func ClearLine() {
	// 创建并打印包含光标移动到行首和清除行的控制字符序列
	fmt.Print("\r\033[K")
}

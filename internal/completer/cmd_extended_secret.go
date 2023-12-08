package completer

import (
	"encoding/json"
	"fmt"
	"github.com/cxweilai/kubectlx/internal/command"
	"github.com/cxweilai/kubectlx/internal/ctx"
	"github.com/cxweilai/kubectlx/internal/kubecli"
	"sigs.k8s.io/yaml"
)

func NewExtendedSecretCommand() *command.Command {
	return &command.Command{
		Name:        "secret",
		Description: "Secret扩展命令，查看明文的Secret",
		Commands: []*command.Command{
			{
				Name:        "show",
				Description: "查询Secret资源详情并明文显示data",
				DynamicParam: &command.DynamicParam{
					Func: func(input string) []*command.Param {
						return kubecli.GetK8sResource("", "v1", "secrets",
							ctx.GetNamespace(), input, LIMIT_SUGGEST)
					},
					Flag:        "SECRET_NAME",
					Description: "查询Secret资源详情并明文显示data",
				},
				Run: WarpHelp(func(cmd *command.ExecCmd) {
					execShowSecret(cmd)
				}),
				Options: []*command.Option{
					{
						Name:        "-o",
						Description: "输出格式，支持:json|yaml，默认是yaml。例如：-o json",
					},
					{
						Name:        "--show-data=true",
						Description: "只显示data",
					},
				},
			},
		},
	}
}

func execShowSecret(cmd *command.ExecCmd) {
	secret, err := kubecli.GetSecretAndBase64Decode(ctx.GetNamespace(), cmd.Param)
	if err != nil {
		fmt.Println(fmt.Sprintf("error: %s", err.Error()))
		return
	}
	if _, ok := cmd.GetOption("--show-data=true"); ok {
		secret = secret["data"].(map[string]interface{})
	}
	format := "yaml"
	if v, ok := cmd.GetOption("-o"); ok {
		format = v
	}
	switch format {
	case "yaml":
		if content, err := yaml.Marshal(secret); err != nil {
			fmt.Println(fmt.Sprintf("error: %s", err.Error()))
		} else {
			fmt.Println(string(content))
		}
	case "json":
		if content, err := json.Marshal(secret); err != nil {
			fmt.Println(fmt.Sprintf("error: %s", err.Error()))
		} else {
			fmt.Println(string(content))
		}
	default:
		fmt.Println(fmt.Sprintf("error: error output foramt %s", format))
	}
}

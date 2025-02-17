package command

import "github.com/spf13/cobra"

var Host string
var hostCommand = &cobra.Command{
	Use:     "go run main.go [-h] 127.0.0.1",
	Short:   "用来设置启动时ip",
	Long:    "用来设置启动时ip",
	Version: "1.0",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			Host = args[0]
		}
		return nil
	},
}

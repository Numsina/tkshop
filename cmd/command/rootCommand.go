package command

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:     "go run [command]",
	Short:   "This is a simple Tiktok mall project, is byte youth training camp's big job",
	Long:    "This is a simple Tiktok mall project, is byte youth training camp's big job",
	Version: "1.0",
}

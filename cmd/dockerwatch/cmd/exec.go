package cmd

import (
	"github.com/gbraad/dockerwatch/pkg/dockerwatch"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(execCmd)
}

var execCmd = &cobra.Command{
	Use:   "exec [-- COMMAND]",
	Short: "execute command",
	Long:  "Execute command",
	Run: func(cmd *cobra.Command, args []string) {
		runExec(args)
	},
}

func runExec(arguments []string) {
	MainLoop(arguments, dockerwatch.Execute)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(execCmd)
}

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "execute command",
	Long:  "Execute command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi")
	},
}

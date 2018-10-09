package cmd

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"github.com/gbraad/dockerwatch/pkg/dockerwatch"
	"github.com/spf13/cobra"
	"os"
)

const (
	descriptionShort = "Execute commands on new containers"
	descriptionLong  = "Execute commands on new containers that match the filter criteria"
	intervalUsage    = "Interval of the watch"
	intervalDefault  = 1
	endpointUsage    = "The host to connect to"
	endpointDefault  = "unix:///var/run/docker.sock"
	filterUsage      = "Filter"
	filterDefault    = "status=running"
)

var (
	interval int
	endpoint string
	filter   string
)

var rootCmd = &cobra.Command{
	Use:   commandName,
	Short: descriptionShort,
	Long:  descriptionLong,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		runPrerun()
	},
	Run: func(cmd *cobra.Command, args []string) {
		runRoot()
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&interval, "interval", "n", intervalDefault, intervalUsage)
	rootCmd.PersistentFlags().StringVarP(&endpoint, "host", "H", endpointDefault, endpointUsage)
	rootCmd.PersistentFlags().StringVarP(&filter, "filter", "f", filterDefault, filterUsage)
}

func runPrerun() {
	fmt.Println(commandName)
}

func runRoot() {
	fmt.Println("No command given")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("ERR:", err.Error())
		os.Exit(1)
	}
}

func MainLoop(arguments []string, action func(docker.Client, string, []string) error) {
	dockerwatch.MainLoop(endpoint, filter, interval, action, arguments)
}

package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/fsouza/go-dockerclient"
	//commands "github.com/gbraad/dockerwatch"
	"github.com/spf13/cobra"
)

const (
	name             = "dockerwatch"
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
	Use:   name,
	Short: descriptionShort,
	Long:  descriptionLong,
	Run: func(cmd *cobra.Command, args []string) {
		// do stuff
	},
}

func Execute() {
	rootCmd.Flags().IntVarP(&interval, "interval", "n", intervalDefault, intervalUsage)
	rootCmd.Flags().StringVarP(&endpoint, "host", "H", endpointDefault, endpointUsage)
	rootCmd.Flags().StringVarP(&filter, "filter", "f", filterDefault, filterUsage)

	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}

	conIDs := []string{}
	listOptions := docker.ListContainersOptions{
		All: true,
	}

	if filter != "" {
		fmt.Println("Filtering:", filter)
		parts := strings.SplitN(filter, "=", 2)
		if len(parts) < 2 || parts[1] == "" {
			// move to Validate of command?
			panic("Invalid filter defined")
		}
		listOptions.Filters = map[string][]string{parts[0]: {parts[1]}}
	}

	for true {

		cons, err := client.ListContainers(listOptions)
		if err != nil {
			panic(err)
		}
		for _, con := range cons {
			if Index(conIDs, con.ID) < 0 {
				conIDs = append(conIDs, con.ID)
				/*
					if execCmd.Happened() {
						command := "exec"
						err := commands.Execute(*client, con.ID, command)
						if err != nil {
							fmt.Println("Err: ", err)
						}
						fmt.Println("Executed:", con.ID)
					}
				*/
			}
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

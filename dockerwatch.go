package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/akamensky/argparse"
	"github.com/fsouza/go-dockerclient"
	"github.com/gbraad/dockerwatch/commands"
)

const (
	intervalUsage   = "Interval of the watch"
	intervalDefault = 1
	endpointUsage   = "The host to connect to"
	endpointDefault = "unix:///var/run/docker.sock"
	filterUsage     = "Filter"
	filterDefault   = "status=running"
)

var (
	interval *int
	endpoint *string
	filter   *string
)

func main() {
	parser := argparse.NewParser("dockerwatch", "Simple Docker container watcher")
	interval = parser.Int("n", "interval", &argparse.Options{Default: intervalDefault, Required: false, Help: intervalUsage})
	endpoint = parser.String("H", "host", &argparse.Options{Default: endpointDefault, Required: false, Help: endpointUsage})
	filter = parser.String("f", "filter", &argparse.Options{Default: filterDefault, Required: false, Help: filterUsage})
	execCmd := parser.NewCommand("exec", "Execute a command on a new container")

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	client, err := docker.NewClient(*endpoint)
	if err != nil {
		panic(err)
	}

	conIDs := []string{}
	listOptions := docker.ListContainersOptions{
		All: false,
	}

	if *filter != "" {
		fmt.Println("Filtering:", *filter)
		parts := strings.SplitN(*filter, "=", 2)
		if len(parts) < 2 || parts[1] == "" {
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

				if execCmd.Happened() {
					err := commands.Execute(*client, con.ID)
					if err != nil {
						fmt.Println("Err: ", err)
					}
					fmt.Println("Executed:", con.ID)
				}
			}
		}

		time.Sleep(time.Duration(*interval) * time.Second)
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

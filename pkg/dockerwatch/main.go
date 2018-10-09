package dockerwatch

import (
	"fmt"
	"strings"
	"time"

	"github.com/fsouza/go-dockerclient"
)

func MainLoop(endpoint string, filter string, interval int, action func(docker.Client, string, []string) error, arguments []string) {
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
				action(*client, con.ID, arguments)
				if err != nil {
					fmt.Println("Err: ", err)
				}
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

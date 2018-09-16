package main

import (
	"fmt"
	"os"
	"time"

	"github.com/akamensky/argparse"
	"github.com/fsouza/go-dockerclient"
)

func main() {
	parser := argparse.NewParser("dockerwatch", "Simple Docker container watcher")
	intervalUsage := "Interval of the watch"
	intervalDefault := 1
	interval := parser.Int("n", "interval", &argparse.Options{Default: intervalDefault, Required: false, Help: intervalUsage})
	endpointUsage := "The host to connect to"
	endpointDefault := "unix:///var/run/docker.sock"
	endpoint := parser.String("H", "host", &argparse.Options{Default: endpointDefault, Required: false, Help: endpointUsage})
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

	for true {

		cons, err := client.ListContainers(docker.ListContainersOptions{All: false})
		if err != nil {
			panic(err)
		}
		for _, con := range cons {
			if Index(conIDs, con.ID) < 0 {
				conIDs = append(conIDs, con.ID)
				if execCmd.Happened() {
					err := Execute(*client, con.ID)
					if err != nil {
						fmt.Println("Err: ", err)
					}
				}
				//fmt.Println("ID: ", con.ID)
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

func Execute(client docker.Client, containerID string) error {
	execConfig := docker.CreateExecOptions{
		Container:    containerID,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: false,
		Tty:          false,
		Cmd:          []string{"touch", "/tmp/file"},
		User:         "root",
	}
	execObj, err := client.CreateExec(execConfig)

	client.StartExecNonBlocking(execObj.ID, docker.StartExecOptions{Detach: true})

	return err
}

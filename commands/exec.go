package commands

import (
	"github.com/fsouza/go-dockerclient"
)

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

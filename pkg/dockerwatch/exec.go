package dockerwatch

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
)

func Execute(client docker.Client, containerID string, arguments []string) error {
	execConfig := docker.CreateExecOptions{
		Container:    containerID,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: false,
		Tty:          false,
		Cmd:          arguments,
		User:         "root",
	}
	execObj, err := client.CreateExec(execConfig)

	client.StartExecNonBlocking(execObj.ID, docker.StartExecOptions{Detach: true})

	fmt.Println("exec:", containerID)

	return err
}

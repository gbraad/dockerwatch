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
		AttachStderr: true,
		Tty:          true,
		Cmd:          arguments,
		User:         "root",
	}

	execObj, err := client.CreateExec(execConfig)

	if err != nil {
		return err
	}

	err = client.StartExec(execObj.ID, docker.StartExecOptions{Detach: true})

	if err != nil {
		return err
	}

	fmt.Println("exec:", containerID)
	return nil
}

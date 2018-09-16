package main

import (
	"fmt"
	"time"

	"github.com/fsouza/go-dockerclient"
)

func main() {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}

	imgIDs := []string{}

	for true {

		imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
		if err != nil {
			panic(err)
		}
		for _, img := range imgs {
			if Index(imgIDs, img.ID) < 0 {
				imgIDs = append(imgIDs, img.ID)
				fmt.Println("ID: ", img.ID)
			}
		}

		time.Sleep(1 * time.Second)
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

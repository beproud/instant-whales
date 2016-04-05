package main

import (
	"fmt"
	"time"

	"github.com/fsouza/go-dockerclient"
)



func listContainers() {
    client, _ := docker.NewClientFromEnv()
    imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
    for _, img := range imgs {
        fmt.Println("ID: ", img.ID)
        fmt.Println("RepoTags: ", img.RepoTags)
        fmt.Println("Created: ", img.Created)
        fmt.Println("Size: ", img.Size)
        fmt.Println("VirtualSize: ", img.VirtualSize)
        fmt.Println("ParentId: ", img.ParentID)
    }
}

func runContainer(image string) (string, error) {
	client, _ := docker.NewClientFromEnv()
	config := docker.Config{
		Image: image,
	}
	c, err := client.CreateContainer(docker.CreateContainerOptions{
		Config: &config,
	})
	if err != nil {
		return "", err
	}
	client.StartContainer(c.ID, &docker.HostConfig{})
	return c.ID, nil
}


func killContainer(id string) error {
	client, _ := docker.NewClientFromEnv()
	client.KillContainer(docker.KillContainerOptions{ID: id})
	return client.RemoveContainer(docker.RemoveContainerOptions{ID: id})
}


func timeoutKill(id string, timeout int) {
	time.Sleep(time.Second * time.Duration(timeout))
	killContainer(id)
}

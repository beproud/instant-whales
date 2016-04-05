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

func runContainer() string {
	client, _ := docker.NewClientFromEnv()
	config := docker.Config{
		Image: "redis",
	}
	c, _ := client.CreateContainer(docker.CreateContainerOptions{
		Config: &config,
	})
	client.StartContainer(c.ID, &docker.HostConfig{})
	return c.ID
}


func killContainer(id string) {
	client, _ := docker.NewClientFromEnv()
	client.KillContainer(docker.KillContainerOptions{ID: id})
	client.RemoveContainer(docker.RemoveContainerOptions{ID: id})
}


func timeoutKill(id string, timeout int) {
	time.Sleep(time.Second * time.Duration(timeout))
	killContainer(id)
}

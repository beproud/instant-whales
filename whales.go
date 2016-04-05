package main

import (
	"sort"
	"time"

	"github.com/fsouza/go-dockerclient"
)



func listImages() []string {
	var r []string
	client, _ := docker.NewClientFromEnv()
	imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
	for _, img := range imgs {
		if img.RepoTags[0] != "<none>:<none>" {
			r = append(r, img.RepoTags[0])
		}
	}
	sort.Strings(r)
	return r
}

func runContainer(image string) (string, error) {
	client, _ := docker.NewClientFromEnv()
	config := docker.Config{
		Image: image,
	}
	host := docker.HostConfig{
		PublishAllPorts: true,
	}
	c, err := client.CreateContainer(docker.CreateContainerOptions{
		Config: &config,
		HostConfig: &host,
	})
	if err != nil {
		return "", err
	}
	client.StartContainer(c.ID, &host)
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

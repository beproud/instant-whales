package main

import (
	"sort"
	"time"

	"github.com/fsouza/go-dockerclient"
)


type Ports map[docker.Port][]docker.PortBinding


type ContainerPs struct {
	ID string `json:"containerId"`
	Image string `json:"image"`
	Status string `json:"status"`
	SizeRw int64 `json:"sizeRw"`
	SizeRootFs int64 `json:"sizeRootFs"`
}

type ContainerInfo struct {
	ID string
	Port string
	Ports Ports
}


// Extract port as string from Ports
func portsToPort(ports Ports) string {
	for _, v := range ports {
		return v[0].HostPort
	}
	return ""
}


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


func listContainers() []ContainerPs {
	var r []ContainerPs
	client, _ := docker.NewClientFromEnv()
	l, _ := client.ListContainers(docker.ListContainersOptions{All: false})
	for _, c := range l {
		r = append(r, ContainerPs{
			ID: c.ID,
			Image: c.Image,
			Status: c.Status,
			SizeRw: c.SizeRw,
			SizeRootFs: c.SizeRootFs})
	}
	return r
}


func runContainer(image string) (ContainerInfo, error) {
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
		return ContainerInfo{}, err
	}
	client.StartContainer(c.ID, &host)
	c, _ = client.InspectContainer(c.ID)
	return ContainerInfo{
		c.ID, portsToPort(c.NetworkSettings.Ports), c.NetworkSettings.Ports,
	}, nil
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

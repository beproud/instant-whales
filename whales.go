package main

import (
	"sort"
	"time"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"github.com/docker/engine-api/types/network"
	"github.com/docker/go-connections/nat"
	"golang.org/x/net/context"
)


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
	Ports nat.PortMap
}


// Extract port as string from Ports
func portsToPort(ports nat.PortMap) string {
	for _, v := range ports {
		return v[0].HostPort
	}
	return ""
}


func listImages() []string {
	var r []string
	c, _ := client.NewEnvClient()
	imgs, _ := c.ImageList(
		context.Background(),
		types.ImageListOptions{All: false},
	)
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
	c, _ := client.NewEnvClient()
	l, _ := c.ContainerList(
		context.Background(),
		types.ContainerListOptions{All: false},
	)
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


func runContainer(image string, memory int) (ContainerInfo, error) {
	cl, _ := client.NewEnvClient()
	mr := int64(memory) * 1024 * 1024
	config := container.Config{
		Image: image,
	}
	host := container.HostConfig{
		PublishAllPorts: true,
		Resources: container.Resources{
			Memory: mr,
			KernelMemory: 16 * 1024 * 1024,
		},
	}
	ctx := context.Background()
	c, err := cl.ContainerCreate(
		ctx,
		&config,
		&host,
		&network.NetworkingConfig{},
		"",
	)
	if err != nil {
		return ContainerInfo{}, err
	}
	cl.ContainerStart(ctx, c.ID)
	ins, _ := cl.ContainerInspect(ctx, c.ID)
	return ContainerInfo{
		ins.ID, portsToPort(ins.NetworkSettings.Ports), ins.NetworkSettings.Ports,
	}, nil
}


func killContainer(id string) error {
	cl, _ := client.NewEnvClient()
	return cl.ContainerKill(context.Background(), id, "SIGKILL")
}


func timeoutKill(id string, timeout int) {
	time.Sleep(time.Second * time.Duration(timeout))
	killContainer(id)
}

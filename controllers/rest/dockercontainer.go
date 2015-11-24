package rest

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/fsouza/go-dockerclient"
)

//DockerContainerRest ...
type DockerContainerRest struct {
	beego.Controller
}

//Get ...
func (This *DockerContainerRest) Get() {

	command := This.Ctx.Input.Params[":command"]
	value := This.Ctx.Input.Params[":value"]
	if value == "" && command != "list" {
		This.CustomAbort(500, "Value must be set!")
	}
	fmt.Printf("Value was %v\n", value)
	client, _ := docker.NewClient(endpoint)
	switch command {
	case "create":
		_, err := client.CreateContainer(docker.CreateContainerOptions{Name: value, Config: &docker.Config{Image: value}})
		if err != nil {
			This.CustomAbort(500, err.Error())
		}
	case "start":
		err := client.StartContainer(value, &docker.HostConfig{})
		if err != nil {
			This.CustomAbort(500, err.Error())
		}
	case "stop":
		err := client.StopContainer(value, 0)
		if err != nil {
			This.CustomAbort(500, err.Error())
		}
	case "delete":
		err := client.RemoveContainer(docker.RemoveContainerOptions{ID: value, Force: true})
		if err != nil {
			This.CustomAbort(500, err.Error())
		}
	case "list":
		containers, _ := client.ListContainers(docker.ListContainersOptions{All: true})
		var ci []ContainerInfo
		for _, k := range containers {
			ci = append(ci, ContainerInfo{
				Name:   strings.TrimLeft(k.Names[0], "/"),
				Status: k.Status,
				Image:  k.Image,
			})
		}
		This.Data["json"] = ci
	default:
		This.CustomAbort(501, "\""+command+"\" is not implemented!")
	}
	This.ServeFormatted()
}

type ContainerInfo struct {
	Name   string
	Status string
	Image  string
}

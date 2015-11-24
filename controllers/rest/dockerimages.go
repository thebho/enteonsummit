package rest

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/fsouza/go-dockerclient"
)

const (
	endpoint = "unix:///var/run/docker.sock"
)

//DockerImageRest ...
type DockerImageRest struct {
	beego.Controller
}

//Get ...
func (This *DockerImageRest) Get() {

	command := This.Ctx.Input.Params[":command"]
	value := This.Ctx.Input.Params[":value"]
	fmt.Printf("Value was %v\n", value)
	switch command {
	case "delete":
		This.Data = deleteImage(value, This)
	case "pull":
		This.Data = pullImage(value, This)
	case "list":
		This.Data = listImages(This)
	default:
		This.CustomAbort(501, "\""+command+"\" is not implemented!")
	}
	This.ServeFormatted()
}

func deleteImage(image string, dir *DockerImageRest) (Data map[interface{}]interface{}) {
	Data = make(map[interface{}]interface{})
	client, _ := docker.NewClient(endpoint)
	err := client.RemoveImage(image)
	if err != nil {
		dir.CustomAbort(500, err.Error())
	} else {
		Data["json"] = "Success!"
	}
	return
}

func pullImage(image string, dir *DockerImageRest) (Data map[interface{}]interface{}) {
	Data = make(map[interface{}]interface{})
	fmt.Printf("Pulling %v\n", image)
	client, _ := docker.NewClient(endpoint)
	fmt.Println("About to pull")
	err := client.PullImage(docker.PullImageOptions{Repository: image}, docker.AuthConfiguration{})
	fmt.Println("Finished")
	if err != nil {
		dir.CustomAbort(500, err.Error())
	} else {
		Data["json"] = "Success!"
	}
	return
}

func listImages(dir *DockerImageRest) (Data map[interface{}]interface{}) {
	Data = make(map[interface{}]interface{})
	fmt.Println("Someone is Listing Images!")
	client, _ := docker.NewClient(endpoint)
	imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
	if err != nil {
		dir.CustomAbort(500, err.Error())
		return
	}
	var Images []ImageData
	for _, k := range imgs {

		Images = append(Images, ImageData{
			Repo:    k.RepoTags[0],
			Size:    fmt.Sprintf("%v", k.Size),
			Created: fmt.Sprintf("%v", k.Created),
			Key:     k.ID,
		})
	}
	Data["json"] = Images
	return
}

//ImageData ...
type ImageData struct {
	Repo    string
	Key     string
	Created string
	Size    string
}

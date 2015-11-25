package rest

import (
	"encoding/json"
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
	fmt.Println("Running Get()")
	data := []byte(`
		[
  {
    "Repo": "java:latest",
    "Key": "cb2addcbc5f44c2abe3aeda527c41d0b3c9793ca0d51795c852bb9e62414600d",
    "Created": "1448008512",
    "Size": "413134"
  },
  {
    "Repo": "docker-repo.gonkulator.io/devinfra:latest",
    "Key": "1c20ff30c7632d4a0923763d97656a22682ff92a31c4318ce551e4ac328f80ab",
    "Created": "1447712050",
    "Size": "0"
  },
  {
    "Repo": "ubuntu:latest",
    "Key": "e9ae3c220b23b699cb5e6914af806219b028c78d5cf6fed8eeca98ffaa8c9b43",
    "Created": "1447115707",
    "Size": "0"
  },
  {
    "Repo": "busybox:latest",
    "Key": "c51f86c283408d1749d066333f7acd5d33b053b003a61ff6a7b36819ddcbc7b7",
    "Created": "1446330175",
    "Size": "0"
  },
  {
    "Repo": "alpine:latest",
    "Key": "8a648f689ddb6408f71f7a41025b56583f5a9a0c5fb6df105d97c8e0ea139910",
    "Created": "1442260874",
    "Size": "5248903"
  }
	]
	`)
	var images []ImageData
	json.Unmarshal(data, &images)
	This.Data["json"] = images

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

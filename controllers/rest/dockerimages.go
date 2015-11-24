package rest

import (
	"github.com/astaxie/beego"
)

//DockerImageRest ...
type DockerImageRest struct {
	beego.Controller
}

//Get ...
func (This *DockerImageRest) Get() {

	command := This.Ctx.Input.Params[":command"]
	value := This.Ctx.Input.Params[":value"]
	switch command {
	case "delete":
		deleteImage(value)
	case "pull":
		pullImage(value)

	default:
		This.Data["json"] = "tobeimplemented"

	}
	This.ServeFormatted()
}

func deleteImage(image string) {

}

func pullImage(image string) {

}

func inspectImage(image string) (Data map[interface{}]interface{}) {

	return
}

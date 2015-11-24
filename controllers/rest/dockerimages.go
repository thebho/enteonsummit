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
	This.Data["json"] = "tobeimplemented"
	This.ServeFormatted()
}

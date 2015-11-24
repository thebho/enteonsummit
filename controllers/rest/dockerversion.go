package rest

import (
	"github.com/astaxie/beego"
)

//DockerVersionRest ...
type DockerVersionRest struct {
	beego.Controller
}

//Get ...
func (This *DockerVersionRest) Get() {
	This.Data["json"] = "tobeimplemented"
	This.ServeFormatted()
}

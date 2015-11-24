package routers

import (
	"enteonsummit/controllers"
	"enteonsummit/controllers/rest"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/rest/docker/image/:command/:value", &rest.DockerImageRest{})
	beego.Router("/rest/echo/:val", &rest.EchoController{})
	beego.Router("/rest/echo", &rest.EchoController{})
	beego.Router("/rest/echo3", &rest.EchoController3{})
	beego.Router("/page3", &controllers.Page3Controller{})
	beego.Router("/images", &controllers.Page2Controller{})

	beego.Router("/", &controllers.IndexController{})
}

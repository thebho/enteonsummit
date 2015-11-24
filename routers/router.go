package routers

import (
	"enteonsummit/controllers"
	"enteonsummit/controllers/rest"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/rest/echo/:val", &rest.EchoController{})
	beego.Router("/rest/echo4/:val", &rest.EchoController4{})
	beego.Router("/rest/echo", &rest.EchoController{})
	beego.Router("/rest/echo3", &rest.EchoController3{})
	beego.Router("/rest/echo4", &rest.EchoController4{})
	beego.Router("/page4", &controllers.Page4Controller{})
	beego.Router("/page3", &controllers.Page3Controller{})
	beego.Router("/page2", &controllers.Page2Controller{})

	beego.Router("/", &controllers.IndexController{})
}

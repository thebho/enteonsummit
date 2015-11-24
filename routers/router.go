package routers

import (
	"github.com/astaxie/beego"
	"edu/Task_06_Beego/task06beego/controllers"
)

func init() {
	beego.Router("/page3", &controllers.Page3Controller{})
  beego.Router("/page2", &controllers.Page2Controller{})
	beego.Router("/", &controllers.IndexController{})
}

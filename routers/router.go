package routers

import (
	"TestWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/getUser", &controllers.GetController{})
	beego.Router("/get", &controllers.ObjectController{})
    beego.Router("/console", &controllers.ConsoleController{})
}

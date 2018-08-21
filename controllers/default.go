package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "kinwe"
	c.Data["Email"] = "huxin304@gmail.com"
	c.TplName = "index.tpl"
}

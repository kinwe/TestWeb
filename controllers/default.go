package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "kinwe"
	c.Data["Email"] = "huxin304@gmail.com"
	c.TplName = "test.html"
}

type GetController struct {
	beego.Controller
}

func (c *GetController) Get() {
	c.TplName = "get.tpl"
}

type ObjectController struct {
	beego.Controller
}
//实现Post方法
func (this *ObjectController) Post() {

	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println(username, password)
	this.TplName = "test.html"
}
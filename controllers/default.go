package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "1"
	c.Data["Email"] = "sdfasf"
	c.TplName = "index.tpl"
}

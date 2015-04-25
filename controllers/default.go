package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["website"] = "角落网"
	c.Data["email"] = "a@nooks.cn"
	c.TplNames = "index.tpl"
	c.Data["html"] = "<div>hello</div>"
}

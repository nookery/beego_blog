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
	c.Data["isHome"] = true
	c.Data["isLogin"] = checkAccout(c.Ctx)
	c.TplNames = "index.tpl"
}

package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	isExit := c.Input().Get("exit") == "true"
	if isExit {
		c.Ctx.SetCookie("username", "", -1, "/")
		c.Ctx.SetCookie("password", "", -1, "/")
		c.Redirect("/", 301)
		return
	}

	c.Data["website"] = "角落网"
	c.Data["email"] = "a@nooks.cn"
	c.TplNames = "login.html"
}

func (c *LoginController) Post() {
	//打印出post的数据
	// c.Ctx.WriteString(fmt.Sprint(c.Input()))
	// return
	username := c.Input().Get("username")
	password := c.Input().Get("password")
	remember := c.Input().Get("remember") == "on"

	if beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password {
		maxAge := 0 //关闭浏览器就清除cookie
		if remember {
			maxAge = 1<<31 - 1 //左移31位
		}
		c.Ctx.SetCookie("username", username, maxAge, "/")
		c.Ctx.SetCookie("password", password, maxAge, "/")
	}

	c.Redirect("/", 301)
	return
}

func checkAccout(ctx *context.Context) bool {
	ck1, err1 := ctx.Request.Cookie("username")
	if err1 != nil {
		return false
	}
	username := ck1.Value

	ck2, err2 := ctx.Request.Cookie("password")
	if err2 != nil {
		return false
	}
	password := ck2.Value

	return beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password
}

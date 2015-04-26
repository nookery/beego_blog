package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	option := c.Input().Get("option")
	switch option {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	case "delete":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	}
	c.Data["website"] = "角落网"
	c.Data["email"] = "a@nooks.cn"
	c.Data["isCategory"] = true
	var err error
	c.Data["categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
		return
	}
	c.TplNames = "category.html"
}

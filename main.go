package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"hello/models"
	_ "hello/routers"
)

func init() {
	models.RegisterDB() // 用来创建数据库
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true) //第二个参数：已存在时是否drop；第三个参数：啰嗦模式
	beego.Run()
}

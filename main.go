package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "hello/routers"
)

func main() {
	beego.Router("/list",&MainController{})
	beego.Run()
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get()  {
	logs.Warn("this is samlai")
	logs.Error("samlai error")
	logs.Info("info msg")
	this.Ctx.WriteString("samlai Hello ...")
}


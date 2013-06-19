package main

import (
	"github.com/astaxie/beego"
	"jethome/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/:name:string", &controllers.MainController{})
	//beego.Router("/contact", &controllers.MainController{})
	beego.Router(`/p/:name([\w]+)`, &controllers.AddController{})
	beego.Router(`/report`, &controllers.ReportController{})
	beego.Info("Jethome started ...")
	beego.Run()
}

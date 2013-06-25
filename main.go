package main

import (
	"github.com/astaxie/beego"
	"jethome/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/:name:string", &controllers.MainController{})
	beego.Router("/p/:pname:string", &controllers.MainController{})
	beego.Router("/report", &controllers.ReportController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/api/:name:string", &controllers.ApiController{})
	beego.Info("started ...")
	beego.Run()
}

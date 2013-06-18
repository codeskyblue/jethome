package main

import (
	"github.com/astaxie/beego"
	"jethome/controllers"
	"jethome/controllers/proj"
	"jethome/controllers/user"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &user.MainController{})
	beego.Router("/proj", &proj.MainController{})
	beego.Router(`/proj/:name([\w]+)`, &proj.AddController{})
	beego.Router(`/report`, &controllers.ReportController{})
	beego.Info("Jethome started ...")
	beego.Run()
}

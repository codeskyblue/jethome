package main

import (
	"github.com/astaxie/beego"
	"jethome/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Info("Jethome started ....")
	beego.Run()
}

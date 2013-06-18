package main

import (
	"github.com/astaxie/beego"
	"jethome/controllers"
	"jethome/controllers/user"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &user.MainController{})
	beego.Info("Jethome started ...")
	beego.Run()
}

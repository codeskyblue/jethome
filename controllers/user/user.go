package user

import (
	"github.com/astaxie/beego"
	"github.com/shxsun/redis"
	"jethome/models"
)

var client redis.Client

func init() {
	client.Addr = "127.0.0.1:6379"
	client.Db = 0
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	users, err := models.ListUser()
	if err != nil {
		beego.Error("list user fail")
		return // TODO: redirect to some where
	}
	this.Data["UserList"] = users
	this.TplNames = "user/list.tpl"
}

func (this *MainController) Post() {
	var user models.User
	user.Name = this.GetString("name")
	user.Email = this.GetString("email")
	models.AddUser(user)
	this.Ctx.Redirect(302, "/user/")
}

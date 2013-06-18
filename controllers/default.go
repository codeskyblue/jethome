package controllers

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
	Info string
}

func (this *MainController) Get() {
    // get proj list from db
	projs, err := models.ListProject()
	if err != nil {
		beego.Error("list proj fail")
	}
	this.Data["ProjList"] = projs

    // TODO: get username from session
	this.Data["Username"] = "astaxie"
	this.Data["Email"] = "astaxie@gmail.com"
    // get  leave message
	if tmp, err := client.Lrange("message", -10, -1); err == nil {
		result := make([]string, 0, 10)
		for _, info := range tmp {
			result = append(result, string(info))
		}
		beego.Trace("all infos are", result)
		this.Data["Infos"] = result
	} else {
		beego.Error(err)
	}
	this.TplNames = "index.tpl"
}

func (this *MainController) Post() {
	info := this.Input().Get("info")
	if info == "" {
		beego.Warn("post got no message")
		this.Ctx.NotFound("no info submit")
	}
	beego.Trace("receive post info:", this.Info)
	this.Data["Username"] = "astaxie"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "submit.tpl"
	client.Lpush("message", []byte(info))
	this.Ctx.Redirect(302, "/")
}

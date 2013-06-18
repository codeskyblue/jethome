package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shxsun/redis"
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
	this.Data["Username"] = "astaxie"
	this.Data["Email"] = "astaxie@gmail.com"
	if tmp, err := client.Lrange("report", 0, -1); err == nil {
		result := make([]string, 0, 20)
		for _, info := range tmp {
			result = append(result, string(info))
		}
		//beego.Trace("all infos are", result)
        beego.Trace("info count:", len(result))
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
	//beego.Trace("receive post info:", this.Info)
	this.Ctx.Redirect(302, "/")
	this.Data["Username"] = "astaxie"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "submit.tpl"
	client.Lpush("report", []byte(info))
}

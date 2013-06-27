// admin.go
package controllers

import (
	"github.com/astaxie/beego"

//	"jethome/models"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {

	content := ""
	this.Data["Admin"] = true
	this.Data["Content"] = content
	this.TplNames = "index.tpl"
}

// TODO: finish post request
func (this *AdminController) Post() {
	name := this.GetString("name")
	qa := this.GetString("qa")
	rd := this.GetString("rd")
	desc := this.GetString("description")
	beego.Debug(name, qa, rd, desc)
	this.Data["Content"] = name
	this.TplNames = "index.tpl"
}

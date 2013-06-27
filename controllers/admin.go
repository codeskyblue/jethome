// admin.go
package controllers

import (
	"github.com/astaxie/beego"
	"jethome/models"
	"strings"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	pm := NameList("")
	this.Data["ProjList"] = pm
	content := ""
	this.Data["Admin"] = true
	this.Data["Content"] = content
	this.TplNames = "index.tpl"
}

func (this *AdminController) Post() {
	var p = models.Project{
		Name:        this.GetString("name"),
		QA:          strings.Fields(this.GetString("qa")),
		RD:          strings.Fields(this.GetString("rd")),
		Description: this.GetString("description"),
	}
	p.Save()
	this.Ctx.Redirect(302, "/")
}

package proj

import (
	"github.com/astaxie/beego"
	"jethome/models"
	"strconv"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	projs, err := models.ListProject()
	if err != nil {
		beego.Error("list proj fail")
		return // TODO: redirect to some where
	}
	this.Data["ProjList"] = projs
	this.TplNames = "proj/index.tpl"
}

func (this *MainController) Post() {
	var pj models.Project
	pj.Name = this.GetString("name")
	pj.Description = this.GetString("description")
	pj.QA = strings.Fields(this.GetString("qa"))
	pj.RD = strings.Fields(this.GetString("rd"))
	pj.Type, _ = strconv.Atoi(this.GetString("type"))
	pj.Save()
	this.Ctx.Redirect(302, "/proj/")
}

type AddController struct {
	beego.Controller
	Job models.Job
}

func (this *AddController) Get() {
	pname := this.Ctx.Params[":name"]
	beego.Trace("router name:", pname)
	if pname == "" {
		this.Ctx.Redirect(302, "/proj/")
	}
	job, err := models.GetJob(pname, 0, -1)
	if err != nil {
		beego.Warn("list jobs error", err)
		this.Ctx.Redirect(302, "/proj")
	}
	this.Data["Jobs"] = job
	this.Job = job
	this.TplNames = "proj/jobs.tpl"
}

func (this *AddController) Post() {
	content := this.GetString("content")
	pname := this.Ctx.Params[":name"]
	if pname == "" {
		this.Ctx.Redirect(302, "/proj/")
	}
	job, err := models.GetJob(pname, 0, -1)
	if err != nil {
		beego.Error("list jobs error", err)
		return
	}
	job.Content = append(job.Content, content)
	err = job.Save()
	if err != nil {
		beego.Error("add job error:", err)
	}
	this.Ctx.Redirect(302, "/proj/"+pname)
}

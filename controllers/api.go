// report.go
package controllers

import (
	"github.com/astaxie/beego"

//	"jethome/models"
)

type ApiController struct {
	beego.Controller
}

func (this *ApiController) Get() {
	args := struct {
		Name string
		Age  int
	}{
		Name: "ssx",
		Age:  12,
	}
	this.Data["json"] = &args
	this.ServeJson()
}

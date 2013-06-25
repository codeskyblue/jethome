// report.go
package controllers

import (
	"github.com/astaxie/beego"
	"jethome/models"
)

type ReportController struct {
	beego.Controller
}

func (r *ReportController) Get() {
	jobs, _ := models.ListJobs()
	quickJobs := make([]models.Job, 0, 20)
	heavyJobs := make([]models.Job, 0, 20)
	for _, job := range jobs {
		switch job.Type {
		case 0:
			quickJobs = append(quickJobs, job)
		case 1:
			heavyJobs = append(heavyJobs, job)
		}
	}
	r.Data["QuickJobs"] = quickJobs
	r.Data["HeavyJobs"] = heavyJobs
	r.TplNames = "report.tpl"
}

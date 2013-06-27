package main

import "jethome/models"
import "github.com/shxsun/beelog"

func main() {
	var p = models.Project{
		Name:        "hermes",
		Description: "shu guang",
	}
	beelog.Info(p)

	err := p.Save()
	beelog.Debug("save error", err)
	beelog.Info(models.Names())
}

package controllers

import (
	"Logmanage/logweb/models"

	"github.com/astaxie/beego"
)

type ModfyController struct {
	beego.Controller
}

func (c *ModfyController) Modfy() {
	f := form{}
	c.ParseForm(&f)
	models.EtcdPut(f.Path, f.Ip, f.Logfile, f.Topic)
	c.Layout = "layout.html"

	c.Redirect("/loglist", 302)

}

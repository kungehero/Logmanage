package controllers

import (
	"Logmanage/logweb/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	Endpoints := []string{"localhost:2379", "localhost:22379", "localhost:32379"}
	c.Layout = "layout.html"

	models.EtcdNew(Endpoints)
	models.EtcdGet("etcd/config")
	c.Data["LogValue"] = models.EtcdValue
	c.TplName = "index.html"
}

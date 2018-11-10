package controllers

import (
	"Logmanage/logweb/models"

	"github.com/astaxie/beego"
)

type AddController struct {
	beego.Controller
}
type form struct {
	Path string `form:"path"`
	Ip   string `form:"ip"`

	Logfile string `form:"logfile"`

	Topic string `form:"topic"`
}

func (c *AddController) Add() {
	f := form{}
	c.ParseForm(&f)
	models.EtcdPut(f.Path, f.Ip, f.Logfile, f.Topic)
	c.Layout = "layout.html"

	c.Redirect("/loglist", 302)

}

func (c *AddController) Delete() {

	key := c.GetString("key")
	c.Layout = "layout.html"

	Endpoints := []string{"localhost:2379", "localhost:22379", "localhost:32379"}

	models.EtcdNew(Endpoints)
	models.EtcdDelete(key)

	c.Redirect("/loglist", 302)

}

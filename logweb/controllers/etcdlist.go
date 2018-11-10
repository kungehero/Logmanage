package controllers

import "github.com/astaxie/beego"

type EaddController struct {
	beego.Controller
}

func (c *EaddController) Get() {

	c.Layout = "layout.html"
	c.TplName = "add.html"
}

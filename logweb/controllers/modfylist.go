package controllers

import "github.com/astaxie/beego"

type EmodfyController struct {
	beego.Controller
}

func (c *EmodfyController) Get() {

	c.Layout = "layout.html"
	c.TplName = "modfy.html"
}

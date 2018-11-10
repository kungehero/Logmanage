package controllers

import (
	"github.com/astaxie/beego"
)

type KafkaController struct {
	beego.Controller
}

func (c *KafkaController) Get() {
	c.Layout = "layout.html"
	c.TplName = "kafka.html"
}

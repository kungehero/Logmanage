package routers

import (
	"Logmanage/logweb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/loglist", &controllers.MainController{})
	beego.Router("/kafkalist", &controllers.KafkaController{})
	beego.Router("/etcdadd", &controllers.EaddController{})
	beego.Router("/modfylist", &controllers.EmodfyController{})

	beego.Router("/add", &controllers.AddController{}, "*:Add")
	beego.Router("/modfy", &controllers.ModfyController{}, "*:Modfy")
	beego.Router("/delete", &controllers.AddController{}, "*:Delete")

}

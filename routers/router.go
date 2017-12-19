package routers

import (
	"jinghui/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/index", &controllers.IndexController{}, "*:Get")
    beego.Router("/item/?:id", &controllers.ItemController{})
	//beego.AutoRouter(&controllers.ItemController{})
}

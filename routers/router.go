package routers

import (
	"github.com/astaxie/beego"
	"jinghui/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.IndexController{}, "*:Get")
	beego.Router("/item/?:id:int", &controllers.ItemController{})
	//beego.AutoRouter(&controllers.ItemController{})
}

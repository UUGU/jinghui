package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type ItemController struct {
	beego.Controller
}

func (c *ItemController) Get() {

	fmt.Printf(c.Ctx.Input.Param(":id"))

	//c.TplName = "index.html"
}

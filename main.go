package main

import (
	"jinghui/models"
	_ "jinghui/routers"

	"github.com/astaxie/beego"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	//commonResponse := common.CommonResponse{}
	//commonResponse.ToSuccess("成功啦！")
	//commonResponse.SetData("独立内容！")
	//commonResponse.SetMapData("key1", "data1")
	//commonResponse.SetMapData(23, 123)
	//logs.Info(commonResponse)
	beego.Run()
}

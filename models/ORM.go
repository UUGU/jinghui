package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "chen:123456@172.25.35.10:3306/quanta?useUnicode=true&characterEncoding=utf8")
}

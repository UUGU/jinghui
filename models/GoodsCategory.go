package models

import (
	"github.com/astaxie/beego/orm"
)

type GoodsCategory struct {
	Id          int64
	JId         int64
	Name        string
	Description string
}

func (g *GoodsCategory) GetGoodsCategory(param GoodsCategory) GoodsCategory{
	o := orm.NewOrm()
	o.Raw("aa")
	return param
}
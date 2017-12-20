package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "chinawsdpay:mQu529Bd376Le2@chinawsdpay.gotoftp1.com:3306/chinawsdpay?useUnicode=true&characterEncoding=utf8")
}

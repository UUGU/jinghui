package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id           int64
	CategoryName string
	Description  string
}

func (m *Category) TableName() string {
	return TableName("category")
}

func GetCategoryById(id int64) (*Category, error) {
	o := orm.NewOrm()
	// o.Raw("aa")
	// ids := []int{1, 2, 3}
	// o.Raw("SELECT * FROM category WHERE id IN (?, ?, ?)", ids)
	category := Category{Id: id}

	ormerr := o.Read(&category)

	if ormerr == orm.ErrNoRows {
		logs.Error("查询不到,主键:", category.Id)
	} else if ormerr == orm.ErrMissPK {
		logs.Error("找不到主键:", category.Id)
	} else {
		return &category, nil
	}
	return nil, ormerr
}

func GetCategories(param Category) ([]Category, error) {
	o := orm.NewOrm()
	// o.Raw("aa")
	// ids := []int{1, 2, 3}
	// o.Raw("SELECT * FROM category WHERE id IN (?, ?, ?)", ids)
	var categorys []Category
	num, err := o.Raw("SELECT * FROM category WHERE category_name = ?", param.CategoryName).QueryRows(&categorys)
	if err != nil {
		return nil, err
	}
	logs.Info("Category nums: ", num)
	//categorys := []*Category{}
	//o.QueryTable(new(Category).TableName()).All(&categorys)
	return categorys, nil
}

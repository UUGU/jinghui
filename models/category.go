package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"jinghui/common"
)

type Category struct {
	Id          int64  `json:"id"`
	ParentId    int64  `json:"parentId"`
	Grade       int64  `json:"grade"`
	Name        string `json:"name"`
	Description string `json:"description"` //`json:"description,omitempty"`
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

func GetCategories(param *Category) ([]Category, error) {
	o := orm.NewOrm()
	// o.Raw("aa")
	// ids := []int{1, 2, 3}
	// o.Raw("SELECT * FROM category WHERE id IN (?, ?, ?)", ids)
	var categorys []Category
	num, err := o.Raw("SELECT * FROM category WHERE name = ?", param.Name).QueryRows(&categorys)
	if err != nil {
		return nil, err
	}
	logs.Info("Category nums: ", num)
	//categorys := []*Category{}
	//o.QueryTable(new(Category).TableName()).All(&categorys)
	return categorys, nil
}

func AddCategory(param *Category) (common.CommonResponse) {
	commonResponse := common.CommonResponse{}
	o := orm.NewOrm()
	result, ormerr := o.Insert(param)
	if nil != ormerr {
		logs.Error("Add Category Error:", ormerr)
		return commonResponse.ToFail(ormerr.Error())
	} else {
		return commonResponse.ToSuccess("Add Category Success! Id:" + string(result))
	}
}

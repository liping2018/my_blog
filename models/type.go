// file ：type.go
// auth : Li ping
// date : 2018-09-24
// desc : 文章类型

package models

import "github.com/astaxie/beego/orm"

type Type struct {
	Typeid int    `orm:"column(typeid);auto"  description:"" json:"typeid"`
	Type   string `orm:"column(type);null"  description:"" json:"type"`
}

func (t *Type) TableName() string {
	return "type"
}

func init() {
	orm.RegisterModel(new(Type))
}

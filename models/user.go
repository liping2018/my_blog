// file ：user.go
// auth : Li ping
// date : 2018-09-24
// desc : 用户实体

package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Userid       int       `orm:"column(userid);auto"  description:"" json:"userid"`
	Username     string    `orm:"column(username);null"  description:"" json:"username"`
	Password     string    `orm:"column(password);null"  description:"" json:"password"`
	Jointime     time.Time `orm:"column(jointime);type(datetime);null"  description:"" json:"jointime"`
	Introduction string    `orm:"column(introduction);null"  description:"" json:"introduction"`
	Adminlable   int       `orm:"column(adminlable);null"  description:"" json:"adminlable"`
	Headpath     string    `orm:"column(headpath);null"  description:"" json:"headpath"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

func AddNewUser(username, password, introduction, headpath string, adminlable int) int64 {

	user := &User{
		Username:     username,
		Password:     password,
		Jointime:     time.Now(),
		Introduction: introduction,
		Adminlable:   0,
		Headpath:     headpath,
	}
	o := orm.NewOrm()
	id, err := o.Insert(user)
	if err != nil {
		fmt.Println("添加成员失败")
	}
	return id
}

//file:initDatabase.go
//auth:liping
//date:2018-09-24
//desc:初始化数据库工具

package utils

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	database := GetString("mysql::database")
	username := GetString("mysql::username")
	password := GetString("mysql::password")
	host := GetString("mysql::host")
	port := GetString("mysql::port")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", username+":"+password+"@("+host+":"+port+")/"+database+"?charset=utf8")
}

// file ：reply.go
// auth : Li ping
// date : 2018-09-24
// desc : 回复实体

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Reply struct {
	Replyid   int       `orm:"column(replyid);auto"  description:"" json:"replyid"`
	Userid    int       `orm:"column(userid);null"  description:"" json:"userid"`
	Articleid int       `orm:"column(articleid);null"  description:"" json:"articleid"`
	Replytime time.Time `orm:"column(replytime);type(datetime);null"  description:"" json:"replytime"`
}

func (t *Reply) TableName() string {
	return "reply"
}

func init() {
	orm.RegisterModel(new(Reply))
}

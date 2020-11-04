package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type Users struct {
	Id       int
	Name     string
	Password string
}

func init() {
	orm.RegisterModel(new(Users))
}

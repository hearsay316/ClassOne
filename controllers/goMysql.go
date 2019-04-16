package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
)

type MysqlController struct {
	beego.Controller
}

func (C *MysqlController) ShowMysql() {
	sql.Open("mysql")
}

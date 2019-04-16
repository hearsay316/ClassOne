package controllers

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlController struct {
	beego.Controller
}

func (C *MysqlController) ShowMysql() {
	conn, err := sql.Open("mysql", "root:qq123456@tcp(sh-cdb-5pd122qv.sql.tencentcdb.com:63231)/classOne?charset=utf8")
	if err != nil {
		fmt.Println("连接错误", err)
		return
	}
	defer conn.Close()
	/*	_,err=conn.Exec("create table userInfo(id int,name varchar(11))")
		if err!=nil{
			fmt.Println("创建错误",err)
			return
		}*/

	/* _,err=conn.Exec(`insert userInfo(id,name)VALUE (?,?)`,1,"title")
	if err!=nil{
		fmt.Println("创建错误",err)
		return
	}*/
	rows, err := conn.Query("select name from userInfo")
	if err != nil {
		fmt.Println("创建错误", err)
		return
	}
	var name string
	for rows.Next() {
		_ = rows.Scan(&name)
		logs.Info(name, 11111111111111)
	}

	C.Ctx.WriteString("sws")
}

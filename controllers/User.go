package controllers

import (
	"classOne/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"net/http"
)

type RegController struct {
	beego.Controller
}

func (C *RegController) ShowReg() {
	C.TplName = "register.html"
}
func (C *RegController) HandleReg() {
	name := C.GetString("userName")
	password := C.GetString("password")
	logs.Info(name, 5555, password)

	if name == "" || password == "" {
		C.TplName = "register.html"
		return
	}
	// 插入数据库
	o := orm.NewOrm()
	user := models.User{UserName: name, PassWord: password}
	_, err := o.Insert(&user)
	if err != nil {
		logs.Info("插入失败")
		C.TplName = "register.html"
		return
	}
	C.Redirect("/", http.StatusFound)

}

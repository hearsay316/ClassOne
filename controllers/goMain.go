package controllers

import (
	"classOne/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (C *MainController) Get() {
	C.TplName = "login.html"
}
func (C *MainController) Post() {
	name := C.GetString("userName")
	passw := C.GetString("password")
	logs.Info(name, 111, passw)
	if name == "" || passw == "" {
		C.TplName = "register.html"
		return
	}
	o := orm.NewOrm()
	user := models.User{UserName: name}
	err := o.Read(&user, "UserName")
	if err != nil {
		logs.Info("用户名失败")
		C.TplName = "login.html"
		return
	}
	if user.PassWord != passw {
		logs.Info("密码失败")
		C.TplName = "login.html"
		return
	}
	C.Ctx.WriteString("登录成功")
}

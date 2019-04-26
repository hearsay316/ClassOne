package controllers

import (
	"classOne/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"net/http"
	"time"
)

type MainController struct {
	beego.Controller
}

func (C *MainController) Get() {
	name := C.Ctx.GetCookie("userName")
	if name == "" {
		C.TplName = "login.html"
	}
	C.Data["Username"] = name

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
	logs.Info(user)
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
	check := C.GetString("remember")
	if check == "on" {
		C.Ctx.SetCookie("userName", name, time.Second*3600)
	} else {
		C.Ctx.SetCookie("userName", "", -1)
	}
	logs.Info("zhge shi ")
	C.SetSession("UserName", name)
	C.Redirect("/Article/ShowArticle", http.StatusFound)
}

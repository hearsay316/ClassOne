package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type IndexController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.comdd"
	c.Data["test"] = "区块链"
	//c.TplName = "index.tpl"
	c.TplName = "test.html"
}
func (this *IndexController) Post() {
	this.Data["test"] = "区块链最棒"
	this.TplName = "test.html"
}
func (this *IndexController) ShowGet() {
	this.Data["test"] = "Get高级路由"
	this.TplName = "test.html"
}
func (this *IndexController) ShowPost() {
	this.Data["test"] = "Post高级路由"
	this.TplName = "test.html"
}

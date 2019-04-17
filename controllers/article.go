package controllers

import "github.com/astaxie/beego"

type ArticleController struct {
	beego.Controller
}

func (C *ArticleController) ShowArticleList() {
	C.TplName = "index.html"
}

func (C *ArticleController) ShowAddArticle() {
	C.TplName = "add.html"
}

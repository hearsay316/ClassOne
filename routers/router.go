package routers

import (
	"classOne/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/Article/*", beego.BeforeRouter, FiltFunc)
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegController{}, "get:ShowReg;post:HandleReg")
	beego.Router("/Article/ShowArticle", &controllers.ArticleController{}, "get:ShowArticleList;post:HandleArticleList")
	beego.Router("/Article/AddArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")
	beego.Router("/Article/showArticleDetail", &controllers.ArticleController{}, "get:ShowArticleDetail")
	beego.Router("/Article/DeleteArticleType", &controllers.ArticleController{}, "get:HandleDeleteType")
	beego.Router("/Article/DeleteArticle", &controllers.ArticleController{}, "get:HandleDelete")
	beego.Router("/Article/UpdateArticle", &controllers.ArticleController{}, "get:HandleGetUpdate;post:HandlePostUpdate")
	beego.Router("/Article/AddArticleType", &controllers.ArticleController{}, "get:HandleGetAddType;post:HandlePostAddType")
	beego.Router("/Logout", &controllers.ArticleController{}, "get:ShowLogout")
}

var FiltFunc = func(ctx *context.Context) {
	Username := ctx.Input.Session("UserName")
	if Username == nil {
		ctx.Redirect(302, "/")
	}
}

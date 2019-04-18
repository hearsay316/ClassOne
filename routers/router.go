package routers

import (
	"classOne/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegController{}, "get:ShowReg;post:HandleReg")
	beego.Router("/ShowArticle", &controllers.ArticleController{}, "get:ShowArticleList")
	beego.Router("/AddArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")
	beego.Router("/showArticleDetail", &controllers.ArticleController{}, "get:ShowArticleDetail")
	beego.Router("/DeleteArticle", &controllers.ArticleController{}, "get:HandleDelete")
	beego.Router("/UpdateArticle", &controllers.ArticleController{}, "get:HandleGetUpdate;post:HandlePostUpdate")

}

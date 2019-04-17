package routers

import (
	"classOne/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/index", &controllers.IndexController{})
	beego.Router("/index", &controllers.IndexController{}, "get:ShowGet;post:ShowPost")
	beego.Router("/mysql", &controllers.MysqlController{}, "get:ShowMysql")
	beego.Router("/orm", &controllers.OrmController{}, "get:ShowMysql")
	beego.Router("/insert", &controllers.OrmController{}, "get:ShowInsert")
	beego.Router("/update", &controllers.OrmController{}, "get:ShowUpdate")

}

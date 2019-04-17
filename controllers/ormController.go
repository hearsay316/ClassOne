package controllers

import (
	"classOne/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type OrmController struct {
	beego.Controller
}
type User struct {
	Id   int64
	Name string
}

// 数据库连接

func (C *OrmController) ShowMysql() {
	//1. 链接数据库
	_ = orm.RegisterDataBase("default", "mysql", "root:qq123456@tcp(sh-cdb-5pd122qv.sql.tencentcdb.com:63231)/classOne?charset=utf8")
	// 2 .注册表
	orm.RegisterModel(new(User))
	// 3.生成表
	_ = orm.RunSyncdb("default", false, true)
	C.Ctx.WriteString("创建成功1")
}
func (C *OrmController) ShowInsert() {

	/*	//1 获取链接对象,orm
		o:=orm.NewOrm()
		// 2插入对象
		user:=User{Name:"张飞"}
		// 插入对象
		_,err:=o.Insert(&user)
		if err!=nil{
			logs.Info("插入失败")
			C.Ctx.WriteString("插入失败")
			return
		}*/

	// 查询对象
	o := orm.NewOrm()
	user := models.User{Id: 1}
	err := o.Read(&user)
	if err != nil {
		logs.Info("查询失败")
		C.Ctx.WriteString("查询失败")
		return
	}
	logs.Info(user)
	C.Ctx.WriteString("查询使用")
}
func (C *OrmController) ShowUpdate() {
	o := orm.NewOrm()
	user := models.User{Id: 1}
	err := o.Read(&user)
	if err != nil {
		logs.Info("查询失败")
		C.Ctx.WriteString("查询失败")
		return
	}
	user.Name = "关羽"
	_, err = o.Update(&user)
	if err != nil {
		logs.Info("更新失败")
		C.Ctx.WriteString("查询失败")
		return
	}
	C.Ctx.WriteString("更新成功")
}

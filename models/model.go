package models

// 1 创建结构体
// 2 初始化语句
import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id   int64
	Name string
}

func init() {
	//1. 链接数据库
	_ = orm.RegisterDataBase("default", "mysql", "root:qq123456@tcp(sh-cdb-5pd122qv.sql.tencentcdb.com:63231)/classOne?charset=utf8")
	// 2 .注册表
	orm.RegisterModel(new(User))
	// 3.生成表
	err := orm.RunSyncdb("default", false, false)
	if err != nil {
		logs.Info("查询失败")
	}
}

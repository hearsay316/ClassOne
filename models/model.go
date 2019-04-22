package models

// 1 创建结构体
// 2 初始化语句
import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id       int64
	UserName string
	PassWord string
}
type Article struct {
	Id          int          `orm:"pk;auto"`
	Title       string       `orm:"size(20)"`
	Content     string       `orm:"size(500)"`
	Img         string       `orm:"size(50);null"`
	Time        time.Time    `orm:"type(datetime);auto_now_add"`
	Count       int          `orm:"default(0)"`
	ArticleType *ArticleType `orm:"rel(fk)"`
}
type ArticleType struct {
	Id       int
	TypeName string     `orm:"size(20)"`
	Articles []*Article `orm:"reverse(many)"`
}

func init() {
	//1. 链接数据库
	_ = orm.RegisterDataBase("default", "mysql", "root:qq123456@tcp(sh-cdb-5pd122qv.sql.tencentcdb.com:63231)/ClassOne?charset=utf8")
	// 2 .注册表
	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	// 3.生成表
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Info("查询失败")
	}
}

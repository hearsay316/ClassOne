package controllers

import (
	"bytes"
	"classOne/models"
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"go-common/library/cache/redis"
	"math"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

type ArticleController struct {
	beego.Controller
}

// 文章列表页
func (C *ArticleController) ShowArticleList() {
	Username := C.GetSession("UserName")
	selects := C.GetString("select")
	if Username == nil {
		C.Redirect("/", 302)
		return
	}

	//1  查询
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	var artList []models.Article
	pageIndex := C.GetString("pageIndex")
	pageIndex1, err := strconv.Atoi(pageIndex)
	if err != nil {
		pageIndex1 = 1
	}
	pageSize := 2
	start := pageSize * (pageIndex1 - 1)
	_, err = qs.Limit(pageSize, start).RelatedSel("ArticleType").All(&artList)
	if err != nil {
		return
	}

	// 根据类型获取数据
	var articleEithType []models.Article

	var count int64
	if !("" != selects) || selects == "0" {
		_, err = qs.Limit(pageSize, start).RelatedSel("ArticleType").All(&articleEithType)
		if err != nil {
			logs.Info("err顶顶顶顶顶顶顶顶顶顶顶顶顶顶顶顶顶顶顶", err)
		}
		count, err = qs.RelatedSel("ArticleType").Count()
		if err != nil {
			return
		}

	} else {
		_, err = qs.Limit(pageSize, start).RelatedSel("ArticleType").Filter("ArticleType__TypeName", selects).All(&articleEithType) // 必须要是article__下划线
		if err != nil {
			logs.Info("err顶顶顶顶顶顶顶顶顶顶顶顶顶顶顶顶顶顶顶", err)
		}
		count, err = qs.RelatedSel("ArticleType").Filter("ArticleType__TypeName", selects).Count()
		if err != nil {
			return
		}
	}
	pageCont := int(math.Ceil(float64(count) / float64(pageSize)))
	// 获取所有数据类型
	var types []models.ArticleType

	/*selects:=C.GetString("select")
	logs.Info(selects)
	o:=orm.NewOrm()
	var articles []models.ArticleType
	o.QueryTable("Article").Filter("ArticleType_Typename",selects).All(&articles)
	logs.Info(articles)*/

	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		logs.Info("redis 数据库错误")
		return
	}
	rel, err := redis.Bytes(conn.Do("get", "types"))
	if err != nil {
		logs.Info("redis.Bytes")
		return
	}
	dec := gob.NewDecoder(bytes.NewReader(rel))
	_ = dec.Decode(&types)
	logs.Info(types, 555)

	if len(types) == 0 {
		_, _ = o.QueryTable("ArticleType").All(&types)
		var buffer bytes.Buffer
		enc := gob.NewEncoder(&buffer)
		_ = enc.Encode(types)
		_, err = conn.Do("set", "types", buffer.Bytes())
		if err != nil {
			logs.Info("redis 数据库错误")
			return
		}
		logs.Info("这个是从数据库")
	}

	logs.Info(articleEithType, 66666666)
	C.Data["selects"] = selects
	C.Data["types"] = types
	C.Data["artList"] = articleEithType
	C.Data["count"] = count
	C.Data["pageCont"] = pageCont
	C.Data["pageIndex"] = pageIndex1
	C.Data["Ss"] = true
	C.Layout = "layout.html"
	C.TplName = "index.html"
}

// 添加文章类型
func (C *ArticleController) ShowAddArticle() {
	o := orm.NewOrm()
	var types []models.ArticleType
	_, err := o.QueryTable("ArticleType").All(&types)
	if err != nil {
		logs.Info("数据库没有类型值")
	}
	C.Data["types"] = types
	C.TplName = "add.html"
}

// 文章类型展示
func (C *ArticleController) HandleArticleList() {

}

/*
  	1 拿数据
	2 判断数据
	3 插入数据库
	4 返回数据
*/
func (C *ArticleController) HandleAddArticle() {
	artName := C.GetString("articleName")
	artContent := C.GetString("content")
	selects := C.GetString("select")
	fileGet, header, err := C.GetFile("uploadname")
	if err != nil {
		logs.Info("err:C.GetFile-")
		return
	}
	defer fileGet.Close()
	ext := path.Ext(header.Filename)
	logs.Info(ext)
	// 获取文件格式
	if ext != ".jpg" && ext != ".png" && ext != ".jpg" {
		logs.Info("格式不正确")
		return
	}

	// 文件大小限制
	if header.Size > 5000000 {
		logs.Info("文件太大")
		return
	}
	// 文件重名  文件名中不可能有: 冒号
	filename := time.Now().Format("2006-01-02 15_04_05")
	err = C.SaveToFile("uploadname", "./static/img/"+filename+ext)
	if err != nil {
		logs.Info("err:SaveToFile-")
		C.Ctx.WriteString("./static/img/" + filename + ext)
		return
	}
	logs.Info(artContent, artName)

	// 插入数据
	//1 获取orm
	var ArticleType models.ArticleType
	ArticleType.TypeName = selects
	o := orm.NewOrm()
	err = o.Read(&ArticleType, "TypeName")
	if err != nil {
		logs.Info("err", err)
		return
	}
	article := models.Article{Title: artName, Content: artContent, Img: "./static/img/" + filename + ext, ArticleType: &ArticleType}
	_, err = o.Insert(&article)
	if err != nil {
		logs.Info("插入数据失败")
		return
	}
	C.Redirect("/Article/ShowArticle", 302)
} // 查看文章详情
func (C *ArticleController) ShowArticleDetail() {
	id, err := C.GetInt("articleId")
	if err != nil {
		return
	}
	o := orm.NewOrm()
	var article models.Article
	article.Id = id
	_ = o.Read(&article)
	article.Count += 1
	_, err = o.Update(&article)
	if err != nil {
		return
	}
	//
	m2m := o.QueryM2M(&article, "Users")
	UserNames := C.GetSession("UserName")
	user := models.User{UserName: UserNames.(string)}
	_ = o.Read(&user, "UserName")
	// 多对多插入
	_, err = m2m.Add(&user)
	if err != nil {
		logs.Info("插入失败")
	}

	//o.LoadRelated(&article, "Users")
	/*err=o.QueryTable("Article").Filter("Users__User__UserName",user).Distinct().Filter("Id",id).One(&article)
	if err!=nil {
		logs.Info("err:article")
	}*/
	var users []models.User
	//err=o.QueryTable("User").Filter("Articles__Article__Id",id).Distinct().All(&users)
	_, err = o.QueryTable("User").Filter("Articles__Article__Id", id).Distinct().All(&users)
	if err != nil {
		logs.Info("QueryTable")
	}
	logs.Info(users)
	C.Data["users"] = users
	C.Data["article"] = article
	C.TplName = "content.html"
}
func (C *ArticleController) HandleDelete() {
	id, err := C.GetInt("id")
	if err != nil {
		logs.Info("C.GetInt()")
		return
	}
	o := orm.NewOrm()
	article := models.Article{Id: id}
	_ = o.Read(&article)
	err = os.Remove(article.Img)
	if err != nil {
		return
	}
	_, err = o.Delete(&article)
	if err != nil {
		return
	}
	C.Redirect("/Article/ShowArticle", http.StatusFound)
}
func (C *ArticleController) HandleGetUpdate() {
	id, err := C.GetInt("id")
	if err != nil {
		return
	}
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err = o.Read(&article)
	if err != nil {
		return
	}
	C.Data["article"] = article
	C.TplName = "update.html"
}
func (C *ArticleController) HandlePostUpdate() {
	id, err := C.GetInt("articleId")
	artName := C.GetString("articleName")
	artContent := C.GetString("content")
	fileGet, header, err := C.GetFile("uploadname")
	if err != nil {
		logs.Info("err:C.GetFile-")
		return
	}
	defer fileGet.Close()
	ext := path.Ext(header.Filename)
	logs.Info(ext)
	// 获取文件格式
	if ext != ".jpg" && ext != ".png" && ext != ".jpg" {
		logs.Info("格式不正确")
		return
	}

	// 文件大小限制
	if header.Size > 5000000 {
		logs.Info("文件太大")
		return
	}
	// 文件重名  文件名中不可能有: 冒号
	filename := time.Now().Format("2006-01-02 15_04_05")
	err = C.SaveToFile("uploadname", "./static/img/"+filename+ext)
	if err != nil {
		logs.Info("err:SaveToFile-")
		C.Ctx.WriteString("./static/img/" + filename + ext)
		return
	}
	logs.Info(artContent, artName)

	//数据处理
	o := orm.NewOrm()
	var article models.Article
	article.Id = id
	err = o.Read(&article)
	if err != nil {
		logs.Info("更新的文章不存在")
		return
	}
	article.Title = artName
	article.Content = artContent
	article.Img = "./static/img/" + filename + ext
	_, _ = o.Update(&article)

	//返回视图
	C.Redirect("/Article/ShowArticle", 302)
}
func (C *ArticleController) HandleGetAddType() {
	o := orm.NewOrm()
	var articleType []models.ArticleType
	_, err := o.QueryTable("ArticleType").All(&articleType)
	C.Data["articleType"] = articleType
	logs.Info(articleType)
	if err != nil {
		logs.Info(err)
	}
	C.Data["types"] = articleType
	C.TplName = "addType.html"
}
func (C *ArticleController) HandlePostAddType() {
	typeName := C.GetString("typeName")
	o := orm.NewOrm()
	var articleType models.ArticleType
	articleType.TypeName = typeName
	_, err := o.Insert(&articleType)
	if err != nil {
		logs.Info("错误")
	}
	C.Redirect("/Article/AddArticleType", 302)
}

// 退出登录
func (C *ArticleController) ShowLogout() {
	C.DelSession("UserName")
	C.Redirect("/", 302)
}

// 删除文章类型
func (C *ArticleController) HandleDeleteType() {
	id := C.GetString("id")
	id2, _ := strconv.Atoi(id)
	if id2 == 0 {
		logs.Info("获取错误")
		return
	}
	o := orm.NewOrm()
	articleType := models.ArticleType{Id: id2}
	_, err := o.Delete(&articleType)
	if err != nil {
		logs.Info("删除错误")
	}
	C.Redirect("/Article/AddArticleType", 302)
}

package controllers

import (
	"classOne/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"net/http"
	"os"
	"path"
	"time"
)

type ArticleController struct {
	beego.Controller
}

// 文章列表页
func (C *ArticleController) ShowArticleList() {
	//1  查询
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	var artList []models.Article
	_, err := qs.All(&artList)
	if err != nil {
		return
	}
	C.Data["artList"] = artList
	C.TplName = "index.html"
}

func (C *ArticleController) ShowAddArticle() {
	C.TplName = "add.html"
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
	o := orm.NewOrm()
	article := models.Article{Title: artName, Content: artContent, Img: "./static/img/" + filename + ext}
	_, err = o.Insert(&article)
	if err != nil {
		logs.Info("插入数据失败")
		return
	}
	C.Redirect("/ShowArticle", 302)
}
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
	C.Redirect("/ShowArticle", http.StatusFound)
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
	C.Redirect("/ShowArticle", 302)
}

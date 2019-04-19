package main

import (
	_ "classOne/models"
	_ "classOne/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	_ = beego.AddFuncMap("PrePage", ShowPrePage)
	_ = beego.AddFuncMap("NextPage", ShowNextPage)
	beego.Run()
}

//后台定义一个函数
func ShowPrePage(pageIndex int) int {
	logs.Info(pageIndex)
	if pageIndex == 1 {
		return pageIndex
	}
	return pageIndex - 1
}

func ShowNextPage(pageIndex int, pageCount int) int {
	if pageIndex == pageCount {
		return pageIndex
	}
	return pageIndex + 1
}

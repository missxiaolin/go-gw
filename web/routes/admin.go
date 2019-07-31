package routes

import (
	"go-gw/bootstrap"
	"go-gw/web/controller/admin"
)

func AdminConfigure(b *bootstrap.Bootstrapper)  {
	// 七牛
	b.GET("/qiniu/get/token", new(admin.QiNiu).GetToken)
	// 分类
	b.POST("/category/add", new(admin.Category).CategoryAdd)
	b.POST("/category/update", new(admin.Category).CategoryUpdate)
	b.GET("/category/info", new(admin.Category).CategoryInfo)
	b.POST("/category/update/status", new(admin.Category).CategoryUpdateStatus)
	b.GET("/category/list", new(admin.Category).GetCategory)
	// 文章
	b.POST("/new/add", new(admin.New).NewAdd)
	b.POST("/new/update", new(admin.New).NewUpdate)
	b.POST("/new/update/status", new(admin.New).NewUpdateStatus)
	b.GET("/new/info", new(admin.New).NewInfo)
	b.GET("/new/list", new(admin.New).NewList)
}

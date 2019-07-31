package routes

import (
	"go-gw/bootstrap"
	"go-gw/web/controller/admin"
)

func AdminConfigure(b *bootstrap.Bootstrapper)  {
	d := b.Group("/admin")
	// 七牛
	d.GET("/qiniu/get/token", new(admin.QiNiu).GetToken)
	// 分类
	d.POST("/category/add", new(admin.Category).CategoryAdd)
	d.POST("/category/update", new(admin.Category).CategoryUpdate)
	d.GET("/category/info", new(admin.Category).CategoryInfo)
	d.POST("/category/update/status", new(admin.Category).CategoryUpdateStatus)
	d.GET("/category/list", new(admin.Category).GetCategory)
	// 文章
	d.POST("/new/add", new(admin.New).NewAdd)
	d.POST("/new/update", new(admin.New).NewUpdate)
	d.POST("/new/update/status", new(admin.New).NewUpdateStatus)
	d.GET("/new/info", new(admin.New).NewInfo)
	d.GET("/new/list", new(admin.New).NewList)
}

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
}

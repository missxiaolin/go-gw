package routes

import (
	"go-gw/bootstrap"
	"go-gw/web/controller"
)

func Configure(b *bootstrap.Bootstrapper)  {
	b.GET("/", new(controller.Index).Welcome)
	// 七牛
	b.GET("/qiniu/get/token", new(controller.QiNiu).GetToken)
	// 分类
	b.GET("/category/add", new(controller.Category).CategoryAdd)
}
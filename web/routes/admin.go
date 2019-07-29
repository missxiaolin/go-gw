package routes

import (
	"go-gw/bootstrap"
	"go-gw/web/controller/admin"
)

func AdminConfigure(b *bootstrap.Bootstrapper)  {
	// 七牛
	b.GET("/qiniu/get/token", new(admin.QiNiu).GetToken)
	// 分类
	b.GET("/category/add", new(admin.Category).CategoryAdd)
}

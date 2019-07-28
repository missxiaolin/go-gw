package routes

import (
	"go-gw/bootstrap"
	"go-gw/web/controller"
)

func Configure(b *bootstrap.Bootstrapper)  {
	b.GET("/", controller.Index)
}
package routes

import (
	"go-gw/bootstrap"
	"go-gw/web/controller/api"
)

func ApiConfigure(b *bootstrap.Bootstrapper)  {
	d := b.Group("/api")
	d.GET("/search/new/list", new(api.New).SearchNewList)
}
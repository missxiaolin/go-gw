package api

import (
	"github.com/gin-gonic/gin"
	"go-gw/web/controller"
	"go-gw/web/form"
	"go-gw/web/formatter"
	"go-gw/web/service"
)

type New struct {
	controller.Base
}

// 对应前端api接口
func (t *New)  SearchNewList(c *gin.Context)  {
	var (
		NewSearchForm form.NewSearchForm
		count         uint
	)
	err := c.BindJSON(&NewSearchForm)
	if err != nil {
		t.Err(c, "json解析错误", 500)
		return
	}
	err = t.Validator(c, NewSearchForm)
	if err != nil {
		t.Err(c, "参错错误", 300)
		return
	}
	NewSearchForm.Status = 2
	newService := new(service.NewService)
	itemList, err := newService.NewList(NewSearchForm)
	if err != nil {
		t.Succ(c, "ok")
		return
	}

	var dataSet = make([]interface{}, 0)

	for _, item := range itemList {
		dataSet = append(dataSet, formatter.NewBase(item))
	}
	newService.GetNewCount(&count)
	t.Succ(c, "ok", gin.H{
		"item":  dataSet,
		"count": count,
	})
}

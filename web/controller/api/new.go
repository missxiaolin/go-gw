package api

import (
	"github.com/gin-gonic/gin"
	"go-gw/common"
	"go-gw/web/controller"
	"go-gw/web/form"
	"go-gw/web/service"
)

type New struct {
	controller.Base
}

// 对应前端api接口
func (t *New) SearchNewList(c *gin.Context)  {
	var (
		NewSearchForm form.NewSearchForm
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

	itemList, err := new(service.NewService).NewList(NewSearchForm)
	if err != nil {
		t.Succ(c, "ok")
		return
	}
	data, err := common.ArrayStructToMap(itemList)
	if err != nil {
		t.Err(c, "解析出错", 500)
		return
	}
	t.Data(c, data)
}

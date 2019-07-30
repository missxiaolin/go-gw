package admin

import (
	"github.com/gin-gonic/gin"
	"go-gw/common"
	"go-gw/web/controller"
	"go-gw/web/form"
	"go-gw/web/service"
)

type Category struct{ controller.Base }

func (t *Category) CategoryAdd(c *gin.Context)  {
	category := &form.CategoryAddForm{
		Name: "ceshi",
		Pid: 1,
		Status: 1,
	}
	t.Validator(c, category)
}

func (t *Category) GetCategory(c *gin.Context)  {
	var (
		categoryFindForm form.CategoryFindForm
	)
	err := c.BindJSON(&categoryFindForm)
	if err != nil {
		t.Err(c, "json解析错误")
		return
	}
	categoryList, err := new(service.CategoryService).Search(categoryFindForm.Name, categoryFindForm.Pid)
	if err != nil {
		t.Succ(c, "ok")
		return
	}
	data, err := common.ArrayStructToMap(categoryList)
	if err != nil {
		t.Err(c, "解析出错")
		return
	}
	t.Data(c, data)
}
package admin

import (
	"github.com/gin-gonic/gin"
	"go-gw/common"
	"go-gw/web/controller"
	"go-gw/web/form"
	"go-gw/web/model"
	"go-gw/web/service"
)

type Category struct{ controller.Base }

// 创建分类
func (t *Category) CategoryAdd(c *gin.Context)  {
	var (
		CategoryAddForm form.CategoryAddForm
		isValidator bool
	)
	err := c.BindJSON(&CategoryAddForm)
	if err != nil {
		t.Err(c, "json解析错误", 500)
		return
	}
	isValidator = t.Validator(c, CategoryAddForm)
	if !isValidator {
		t.Err(c, "参错错误", 300)
		return
	}
	_, err = new(service.CategoryService).CategoryAdd(model.Category{
		Name: CategoryAddForm.Name,
		Pid: CategoryAddForm.Pid,
		Status: CategoryAddForm.Status,
	})

	if err != nil {
		t.Err(c, "创建分类失败", 500)
		return
	}

	t.Succ(c, "ok")



}

// 获取列表
func (t *Category) GetCategory(c *gin.Context)  {
	var (
		categoryFindForm form.CategoryFindForm
	)
	err := c.BindJSON(&categoryFindForm)
	if err != nil {
		t.Err(c, "json解析错误", 500)
		return
	}
	categoryList, err := new(service.CategoryService).Search(categoryFindForm.Name, categoryFindForm.Pid)
	if err != nil {
		t.Succ(c, "ok")
		return
	}
	data, err := common.ArrayStructToMap(categoryList)
	if err != nil {
		t.Err(c, "解析出错", 500)
		return
	}
	t.Data(c, data)
}
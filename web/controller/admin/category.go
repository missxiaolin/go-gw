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
	)
	err := c.BindJSON(&CategoryAddForm)
	if err != nil {
		t.Err(c, "json解析错误", 500)
		return
	}
	err = t.Validator(c, CategoryAddForm)
	if err != nil {
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

// 修改分类
func (t *Category) CategoryUpdate(c *gin.Context) {
	var (
		CategoryUpdateForm form.CategoryUpdateForm
	)
	err := c.BindJSON(&CategoryUpdateForm)
	if err != nil {
		t.Err(c, "json解析错误", 500)
		return
	}
	err = t.Validator(c, CategoryUpdateForm)
	if err != nil {
		t.Err(c, "参错错误", 300)
		return
	}
	categoryService := new(service.CategoryService)
	category, err := categoryService.GetCategoryByID(CategoryUpdateForm.Id)
	if err != nil {
		t.Err(c, err.Error(), 500)
		return
	}
	category.Name = CategoryUpdateForm.Name
	category.Status = CategoryUpdateForm.Status
	category.Pid = CategoryUpdateForm.Pid
	categoryService.CategoryUpdate(category)
	t.Succ(c, "ok")
}

// 查询分类
func (t *Category) CategoryInfo(c *gin.Context) {
	id := c.Query("id")
	if id == "" || id == "0" {
		t.Err(c, "参错错误", 300)
		return
	}

	category, err := new(service.CategoryService).GetCategoryByID(id)
	if err != nil {
		t.Err(c, err.Error(), 500)
		return
	}
	t.Succ(c, "ok", category)
}

// 更改分类状态
func (t *Category) CategoryUpdateStatus(c *gin.Context)  {
	var (
		CategoryUpdateStatusForm form.CategoryUpdateStatusForm
	)
	err := c.BindJSON(&CategoryUpdateStatusForm)
	if err != nil {
		t.Err(c, "json解析错误", 500)
		return
	}
	err = t.Validator(c, CategoryUpdateStatusForm)
	if err != nil {
		t.Err(c, "参错错误", 300)
		return
	}
	categoryService := new(service.CategoryService)
	category, err := categoryService.GetCategoryByID(CategoryUpdateStatusForm.Id)
	if err != nil {
		t.Err(c, err.Error(), 500)
		return
	}
	category.Status = CategoryUpdateStatusForm.Status
	categoryService.CategoryUpdate(category)
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
package admin

import (
	"github.com/gin-gonic/gin"
	"go-gw/web/controller"
	"go-gw/web/form"
	"go-gw/web/formatter"
	"go-gw/web/model"
	"go-gw/web/service"
)

type New struct{ controller.Base }

// 添加文章
func (t *New) NewAdd(c *gin.Context) {
	var (
		NewAddForm form.NewAddForm
		err        error
	)
	err = c.BindJSON(&NewAddForm)
	if err != nil {
		t.Err(c, "json解析错误", 500)
		return
	}
	err = t.Validator(c, NewAddForm)
	if err != nil {
		t.Err(c, "参错错误", 300)
		return
	}
	_, err = new(service.NewService).NewAdd(model.Article{
		Cid:         NewAddForm.Cid,
		Title:       NewAddForm.Title,
		Author:      NewAddForm.Author,
		Content:     NewAddForm.Content,
		Keywords:    NewAddForm.Keywords,
		Description: NewAddForm.Description,
		Status:      NewAddForm.Status,
	})
	if err != nil {
		t.Err(c, err.Error(), 500)
		return
	}

	t.Succ(c, "ok")
}

// 修改文章
func (t *New) NewUpdate(c *gin.Context) {
	var (
		NewUpdateForm form.NewUpdateForm
		err           error
	)
	err = c.BindJSON(&NewUpdateForm)
	if err != nil {
		t.Err(c, "json解析错误", 500)
		return
	}
	err = t.Validator(c, NewUpdateForm)
	if err != nil {
		t.Err(c, "参错错误", 300)
		return
	}

	newService := new(service.NewService)
	article, err := newService.GetNewById(NewUpdateForm.Id)
	if err != nil {
		t.Err(c, err.Error(), 500)
		return
	}
	article.Cid = NewUpdateForm.Cid
	article.Title = NewUpdateForm.Title
	article.Author = NewUpdateForm.Author
	article.Content = NewUpdateForm.Content
	article.Keywords = NewUpdateForm.Keywords
	article.Description = NewUpdateForm.Description
	article.Status = NewUpdateForm.Status
	newService.NewUpdate(article)
	t.Succ(c, "ok")
}

// 获取文章详情
func (t *New) NewInfo(c *gin.Context) {
	id := c.Query("id")
	if id == "" || id == "0" {
		t.Err(c, "参错错误", 300)
		return
	}

	article, err := new(service.NewService).GetNewById(id)
	if err != nil {
		t.Err(c, err.Error(), 500)
		return
	}
	data := formatter.NewBase(article)
	t.Succ(c, "ok", data)
}

// 修改文章状态
func (t *New) NewUpdateStatus(c *gin.Context) {
	var (
		NewUpdateStatusForm form.NewUpdateStatusForm
	)
	err := c.BindJSON(&NewUpdateStatusForm)
	if err != nil {
		t.Err(c, "json解析错误", 500)
		return
	}
	err = t.Validator(c, NewUpdateStatusForm)
	if err != nil {
		t.Err(c, "参错错误", 300)
		return
	}
	newService := new(service.NewService)
	article, err := newService.GetNewById(NewUpdateStatusForm.Id)
	if err != nil {
		t.Err(c, err.Error(), 500)
		return
	}
	article.Status = NewUpdateStatusForm.Status
	newService.NewUpdate(article)
	t.Succ(c, "ok")
}

// 文章列表
func (t *New) NewList(c *gin.Context) {
	var (
		NewSearchForm form.NewSearchForm
		count         uint
	)
	err := c.BindQuery(&NewSearchForm)
	if err != nil {
		t.Err(c, "json解析错误", 500)
		return
	}
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

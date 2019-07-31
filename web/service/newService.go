package service

import (
	"errors"
	"go-gw/lib/database"
	"go-gw/web/form"
	"go-gw/web/model"
)

type NewService struct {
	BaseService
}

// 创建文章
func (s *NewService) NewAdd(article model.Article) (model.Article, error) {
	db := database.NewDB()
	db.Create(&article)
	if db.NewRecord(article) == true {
		return article, errors.New("创建文章失败")
	}
	return article,nil
}

// 获取文章详情
func (s *NewService) GetNewById(id interface{}) (article model.Article, err error) {
	database.NewDB().Where("id = ?", id).First(&article)
	if article.ID == 0 {
		err = errors.New("未找到文章")
	}
	return article, err
}

// 修改文章
func (s *NewService) NewUpdate(new model.Article) {
	database.NewDB().Save(new)
	return
}

// 文章列表
func (s *NewService) NewList(form form.NewSearchForm) (itemList []model.Article, err error) {
	db := database.NewDB()

	if form.Cid != 0 {
		db = db.Where("cid = ?", form.Cid)
	}
	db.Limit(form.PageSize).Offset((form.Page - 1) * form.PageSize).Find(&itemList)

	return
}

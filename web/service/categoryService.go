package service

import (
	"errors"
	"go-gw/lib/database"
	"go-gw/web/model"
)

type CategoryService struct {
	BaseService
}

// 创建分类
func (s *CategoryService) CategoryAdd(category model.Category) (model.Category, error) {
	db := database.NewDB()
	db.Create(&category)
	if db.NewRecord(category) == true {
		return category, errors.New("创建分类失败")
	}
	return category, nil
}

// 根据id查找分类
func (s *CategoryService) GetCategoryByID(id interface{}) (category model.Category, err error) {
	database.NewDB().Where("id = ?", id).First(&category)
	if category.ID == 0 {
		err = errors.New("未查到到分类")
	}
	return category, err
}

// 修改分类
func (s *CategoryService) CategoryUpdate(category model.Category) {
	database.NewDB().Save(category)
}

// 搜索分类
func (s *CategoryService) Search(name string, pid uint) (itemList []model.Category, err error) {
	db := database.NewDB()
	if name != "" {
		db = db.Where("name = ?", name)
	}
	db = db.Where("pid = ?", pid)

	db.Find(&itemList)
	return
}
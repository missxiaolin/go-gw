package service

import (
	"go-gw/lib/database"
	"go-gw/web/model"
)

type CategoryService struct {
	BaseService
}

func (s *CategoryService) Search(name string, pid uint) (itemList []model.Category, err error) {
	db := database.NewDB()
	if name != "" {
		db = db.Where("name = ?", name)
	}
	db = db.Where("pid = ?", pid)

	db.Find(&itemList)
	return
}
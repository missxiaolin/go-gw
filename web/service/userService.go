package service

import (
	"errors"
	"go-gw/common"
	"go-gw/lib/database"
	"go-gw/web/form"
	"go-gw/web/model"
)

type UserService struct {
	BaseService
}

// 用户登录
func (t *UserService) UserLogin(form form.UserLoginForm) (users model.Users, err error) {
	db := database.NewDB().Where("name = ?", form.Name)
	db.First(&users)
	if users.ID == 0 {
		return users, errors.New("用户账号密码错误")
	}
	if users.Password != common.MD5(form.Password) {
		return users, errors.New("用户账号密码错误")
	}
	return users, nil
}
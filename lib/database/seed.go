package database

import (
	"go-gw/common"
	"go-gw/web/model"
)

func Seed()  {
	seedUser()
}

func seedUser()  {
	db := NewDB()
	user := model.Users{
		Name:     "admin",
		Password: common.MD5("123456"),
		Status: 1,
	}
	db.First(&user, user)
	if user.ID == 0 {
		db.Create(&user)
		if db.NewRecord(user) {
			panic("user数据填充失败")
		}
	} else {
		panic("user数据已经填充")
	}
}


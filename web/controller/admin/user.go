package admin

import (
	"github.com/gin-gonic/gin"
	"go-gw/web/controller"
)

type User struct {
	controller.Base
}

func (t *User) Login(c *gin.Context) {

}

// 添加用户信息
func (t *User) UserAdd(c *gin.Context) {

}

// 修改用户信息
func (t *User) UserUpdate(c *gin.Context) {

}

// 修改用户状态
func (t *User) UserUpdateStatus(c *gin.Context) {

}

// 用户详情
func (t *User) UserInfo(c *gin.Context) {

}

// 用户列表
func (t *User) UserList(c *gin.Context) {

}
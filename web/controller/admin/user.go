package admin

import (
	"github.com/gin-gonic/gin"
	"go-gw/config"
	"go-gw/lib/jwt"
	"go-gw/web/controller"
	"go-gw/web/form"
	"go-gw/web/service"
	"time"
)

type User struct {
	controller.Base
}

// 登录
func (t *User) Login(c *gin.Context) {
	var (
		UserLoginForm form.UserLoginForm
	)
	err := c.BindJSON(&UserLoginForm)

	if err != nil {
		t.Err(c, "json解析错误", 500)
		return
	}
	err = t.Validator(c, UserLoginForm)
	if err != nil {
		t.Err(c, "参错错误", 300)
		return
	}
	users, err := new(service.UserService).UserLogin(UserLoginForm)
	if err != nil {
		t.Err(c, err.Error(), 500)
		return
	}
	var data = map[string]interface{}{
		"id":   users.ID,
		"name": users.Name,
	}
	token, err := jwt.Create(data, config.JWT_SECRET_ADMIN, time.Now().Add(time.Hour * config.JWT_EXP_ADMIN).Unix())
	if err != nil {
		t.Err(c, "登陆失败", 500)
		return
	}
	c.Request.Header.Set("token", token)

	t.Succ(c, "ok", token)

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
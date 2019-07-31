package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smokezl/govalidators"
	"go-gw/common"
	"net/http"
)

type Base struct{
	Ctx *gin.Context
}

// 返回
func (t *Base) Api(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

// 成功
func (t *Base) Succ(c *gin.Context, msg string, data ...interface{}) {

	var d interface{}
	if data != nil && len(data) > 0 {
		d = data[0]
	}

	t.Api(c, http.StatusOK, gin.H{
		"success": true,
		"msg": msg,
		"data": d,
	})
}

// 成功
func (t *Base) Data(c *gin.Context, data interface{}) {
	t.Api(c, http.StatusOK, data)
}

// 转换
func (t *Base) Array(c *gin.Context, data interface{}) {

	data, err := common.ArrayStructToMap(data)
	if err != nil {
		t.Err(c, err.Error(), 500)
		return
	}
	t.Data(c, data)
}

// 抛错
func (t *Base) Err(c *gin.Context, msg string, ErrorCode uint) {

	fmt.Println("[Error]", msg)
	t.Api(c, http.StatusOK, gin.H{
		"success": false,
		"ErrorCode": ErrorCode,
		"msg": msg,
	})
}

// 验证器
func (t *Base) Validator(c *gin.Context, ValidatorData interface{}) (err error) {
	validator := govalidators.New()
	err = validator.LazyValidate(ValidatorData)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	return nil
}
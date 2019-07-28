package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gw/common"
	"net/http"
)

type Base struct{
	Ctx *gin.Context
}

func (t *Base) Api(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

func (t *Base) Succ(c *gin.Context, msg string) {
	t.Api(c, http.StatusOK, gin.H{
		"msg": msg,
	})
}

func (t *Base) Data(c *gin.Context, data interface{}) {
	t.Api(c, http.StatusOK, data)
}

func (t *Base) Array(c *gin.Context, data interface{}) {

	data, err := common.ArrayStructToMap(data)
	if err != nil {
		t.Err(c, err.Error())
		return
	}
	t.Data(c, data)
}

func (t *Base) Err(c *gin.Context, msg string) {

	fmt.Println("[Error]", msg)
	t.Api(c, http.StatusBadRequest, gin.H{
		"msg": msg,
	})
}
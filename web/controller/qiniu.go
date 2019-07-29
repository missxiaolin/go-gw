package controller

import (
	"github.com/gin-gonic/gin"
	"go-gw/lib/qiniu"
)

type QiNiu struct{ Base }

func (t *QiNiu) GetToken(c *gin.Context)  {
	token := qiniu.GetToken()
	t.Succ(c, "ok", token)
}
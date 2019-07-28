package controller

import (
	"github.com/gin-gonic/gin"
	"go-gw/common"
)

func Index(c *gin.Context)  {
	common.ResponseSuccess(c)
}

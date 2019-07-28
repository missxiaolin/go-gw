package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type result struct {
	Success      bool
	Data         interface{}
	ErrorCode    uint32
	ErrorMessage string
}

func ResponseSuccess(c *gin.Context, data ...interface{}) {

	var d interface{}
	if data != nil && len(data) > 0 {
		d = data[0]
	}

	c.JSON(http.StatusOK, result{
		Success: true,
		ErrorCode: 0,
		ErrorMessage: "",
		Data: d,
	})
}

func ResponseErr(c *gin.Context, ErrorCode uint32, ErrorMessage string)  {
	c.JSON(http.StatusOK, result{
		Success: false,
		ErrorCode: ErrorCode,
		ErrorMessage: ErrorMessage,
	})
}
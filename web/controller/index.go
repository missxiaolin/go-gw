package controller

import (
	"github.com/gin-gonic/gin"
)

type Index struct{ Base }

// 欢迎
func (t *Index) Welcome(c *gin.Context)  {
	t.Succ(c, "welcome", nil)
}

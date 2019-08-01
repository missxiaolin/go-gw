package middleware

import (
	"github.com/gin-gonic/gin"
	"go-gw/config"
	"go-gw/lib/jwt"
	"net/http"
)

// 登录校验
func AdminAuth(c *gin.Context)  {
	tokenStr := c.Request.Header.Get("TOKEN")
	if jwt.CheckValid(tokenStr, config.JWT_SECRET_ADMIN) == false {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"success": false,
			"ErrorCode": "700",
			"msg": "登陆失败",
		})
		return
	}

	jwtData, err := jwt.ParseInfo(tokenStr, config.JWT_SECRET_ADMIN)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"success": false,
			"ErrorCode": "700",
			"msg": "异常登录信息",
		})
		return
	}
	info, ok := jwtData["info"]
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"success": false,
			"ErrorCode": "700",
			"msg": "异常登录信息",
		})
		return
	}
	c.Set("admin", info.(map[string]interface{}))
	c.Next()
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"temp/handler"
	"temp/util"
	"time"
)

func Parse(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		c.JSON(200, handler.ErrorInfo{Code: 40001,
			Message: "Authorization不能为空"})
		c.Abort()
		return
	}
	
	claims, err := util.Parse(auth)
	if err != nil {
		c.JSON(200, handler.ErrorInfo{Code: 40001,
			Message: "Authorization格式错误"})
		c.Abort()
		return
	}
	
	if claims.ExpiresAt < time.Now().Unix() {
		c.JSON(200, handler.ErrorInfo{Code: 40001,
			Message: "无效的Authorization"})
		c.Abort()
		return
	}
	c.Set("account", claims.Account)
}

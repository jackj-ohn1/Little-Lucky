package route

import (
	"github.com/gin-gonic/gin"
	"temp/handler"
	"temp/internal/middleware"
)

func GenerateRouter() *gin.Engine {
	engine := gin.Default()
	
	// 返回token以及是否进行邮箱绑定的bool值
	engine.POST("/api/v1/login", handler.Login)
	
	// 发送验证码
	engine.GET("/api/v1/email", handler.SendEmail)
	
	// 验证码正确则绑定验证码
	engine.POST("/api/v1/email/bind", middleware.Parse, handler.ValidateEmail)
	
	return engine
}

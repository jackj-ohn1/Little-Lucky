package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"temp/internal/repository/dao"
	"temp/util"
)

var (
	ERRWRONGIDENTIFICATION = &ErrorInfo{Code: 30200, Message: "学号或密码错误"}
	ERRBINDDATA            = &ErrorInfo{Code: 40000, Message: "数据绑定失败"}
	SUCCESS                = &ErrorInfo{Code: 20000, Message: "操作成功"}
)

func Login(c *gin.Context) {
	var identification Identification
	if err := c.ShouldBindWith(&identification, binding.JSON); err != nil {
		c.JSON(200, ERRBINDDATA)
		return
	}
	_, err := util.Login(identification.Id, identification.Password)
	if err != nil {
		c.JSON(200, ERRWRONGIDENTIFICATION)
		return
	}
	var token string
	token, err = util.Generate(identification.Id)
	if err != nil {
		log.Println(err)
		c.JSON(500, ERRSERVERWRONG)
		return
	}
	
	err, exist := dao.Check(identification.Id)
	if err != nil {
		c.JSON(500, ERRSERVERWRONG)
		return
	}
	
	// exist == true代表要跳转至邮箱绑定
	// 否则直接完成登录
	c.JSON(200, gin.H{
		"code":    20000,
		"message": "操作成功",
		"data":    token,
		"email":   exist,
	})
}

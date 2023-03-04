package handler

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"temp/config"
	"temp/internal/repository/dao"
	"time"
)

type Storage interface {
	setStorage(email, code string)
	clearStorage(email string)
	NewInnerStorage()
}

type innerStorage struct {
	storage  map[string]string
	codeChan chan string
}

func NewInnerStorage() *innerStorage {
	var one innerStorage
	one.storage = make(map[string]string)
	one.codeChan = make(chan string)
	
	go func() {
		for {
			select {
			case person := <-one.codeChan:
				delete(one.storage, person)
			}
		}
	}()
	
	return &one
}

func (i *innerStorage) setStorage(email, code string) {
	i.storage[email] = code
}

func (i *innerStorage) clearStorage(email string) {
	for {
		select {
		case <-time.After(time.Minute * 1):
			i.codeChan <- email
			return
		}
	}
}

var (
	// 鉴于验证码的时效性，此处直接存在内存中
	// 如果服务器性能不行，可以改成redis
	storage        = NewInnerStorage()
	ERREMAILSEND   = &ErrorInfo{Code: 50001, Message: "验证码发送失败"}
	ERRVALIDATION  = &ErrorInfo{Code: 40001, Message: "无效验证码"}
	ERRWRONGDATA   = &ErrorInfo{Code: 30005, Message: "数据错误"}
	ERRSERVERWRONG = &ErrorInfo{Code: 50001, Message: "服务器错误"}
)

func SendEmail(c *gin.Context) {
	dialer := gomail.NewDialer(config.Conf.Email.Host, config.Conf.Email.Port, config.Conf.Email.Sender, config.Conf.Email.SecretKey)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	
	email := c.Query("email")
	
	// 前端做发送时间限制
	if err, code := newEmail(dialer, email); err != nil {
		log.Println(err)
		c.JSON(200, ERREMAILSEND)
		return
	} else {
		storage.setStorage(email, code)
		go storage.clearStorage(email)
	}
	
	c.JSON(200, SUCCESS)
}

type Auth struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}

func ValidateEmail(c *gin.Context) {
	var auth Auth
	if err := c.ShouldBindWith(&auth, binding.JSON); err != nil {
		c.JSON(200, ERRBINDDATA)
		return
	}
	if actual, exist := storage.storage[auth.Email]; !exist || auth.Code != actual {
		c.JSON(200, ERRVALIDATION)
		return
	}
	
	account := c.MustGet("account").(string)
	if auth.Email == "" || account == "" {
		c.JSON(200, ERRWRONGDATA)
		return
	}
	if err := dao.Insert(account, auth.Email); err != nil {
		c.JSON(500, ERRSERVERWRONG)
		return
	}
	
	c.JSON(200, SUCCESS)
}

func newEmail(dialer *gomail.Dialer, emailOwner string) (error, string) {
	email := gomail.NewMessage()
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000000))
	
	email.SetHeader("From", config.Conf.Email.Sender)
	email.SetHeader("To", emailOwner)
	email.SetHeader("Subject", "小幸运邮箱认证")
	email.SetBody("text/html", code)
	
	return dialer.DialAndSend(email), code
}

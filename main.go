package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/wilikidi/go_gin_test/controller"
	"github.com/wilikidi/go_gin_test/utils"
	"net/http"
)

func main() {
	utils.Config()
	utils.InitLog()
	// 初始化一个 validator
	binding.Validator = utils.VALIDATE

	// gin 框架注册制中间件
	router := gin.Default()
	// set the cors。when https
	router.Use(func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,accept, content-type, referer, sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, Tagvalue, User-Agent")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	})

	router.GET("/api/test/version", controller.Version)
	router.GET("/api/test/showheader", controller.ShowHeader)
	router.POST("/api/test/common", controller.Common)
	router.POST("/api/test/community", controller.Community)
	router.POST("/api/test/message", controller.GetMessageByName)
	router.POST("/api/test/showdir", controller.ShowDir)
	router.POST("/api/test/ascii_trasnlate", controller.AsciiTranslate)
	router.GET("/api/test/killself", controller.KillSelf)
	router.POST("/api/test/showBody", controller.ShowBody)
	router.POST("/api/test/showForm", controller.ShowForm)
	router.GET("/api/test/health", controller.Health)

	// 公司一组url
	// 用来进行 key 的解谜
	router.POST("/api/tvu/decrypt_key", controller.DecryptSecretKey)
	router.POST("/api/tvu/signature", controller.GetSignature)

	// 生成数据一组api
	router.GET("/api/data/student", controller.GenerateStudent)
	router.GET("/api/data/student/api/health", controller.StudentHealth)

	// apifox test
	router.POST("/api/token/access", controller.Login)
	router.POST("/api/token/check", controller.Check)
	// todo: 缺少 form 表单系列组 api

	//router.RunTLS(":5002", "server.crt", "server.key")
	router.Run(fmt.Sprintf(":%v", utils.HttpPort))
}

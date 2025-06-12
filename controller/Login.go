package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

var (
	tokenIndex = 0
)

type LoginRequest struct {
	Account string `json:"account"`
	Passwd  string `json:"passwd"`
}

type LoginResp struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	var lr LoginRequest

	if err := c.ShouldBindBodyWith(&lr, binding.JSON); err != nil {
		c.String(400, "%v", err)
		return
	}

	var loginResp LoginResp
	loginResp.Token = fmt.Sprintf("abc_wyk_@_9k9k%v", tokenIndex)
	tokenIndex++

	c.JSON(http.StatusOK, loginResp)
}

type CheckRequest struct {
	Token string `json:"token"`
}

func Check(c *gin.Context) {
	var cr CheckRequest
	if err := c.ShouldBindWith(&cr, binding.JSON); err != nil {
		c.String(400, "%v", err)
		return
	}

	fmt.Println(cr.Token)

	token := c.Request.Header.Get("token")
	fmt.Println(token)

	c.String(200, "%v", "success")
}

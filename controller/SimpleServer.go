package controller

import (
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Message struct {
	UserName string `json:"username" binding:"required"`
}

type User struct {
	UserName string `json:"username,omitempty"`
	UserId   string `json:"userId,omitempty"`
	Age      int    `json:"age,omitempty"`
}

func GetMessageByName(c *gin.Context) {
	var m Message
	if err := c.BindJSON(&m); err != nil {

		log.Infof("[GetMessageByName] error is: %v", err)

		c.JSON(http.StatusBadRequest, err)
		return
	}

	u := User{UserName: m.UserName + "1", Age: 10}

	c.JSON(http.StatusOK, u)

}

func GetUserIdByName(c *gin.Context) {
	var m Message
	if err := c.BindJSON(&m); err != nil {

		log.Infof("[GetUserIdByName] error is: %v", err)

		c.JSON(http.StatusBadRequest, err)
		return
	}

	u := User{UserId: m.UserName + "userid"}

	c.JSON(http.StatusOK, u)

}

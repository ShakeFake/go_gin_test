package controller

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/wilikidi/go_gin_test/model"
	"io/ioutil"
	"net/http"
	"os"
)

type Message struct {
	UserName string `json:"username" binding:"required"`
}

type User struct {
	UserName string `json:"username,omitempty"`
	UserId   string `json:"userId,omitempty"`
	Age      int    `json:"age,omitempty"`
}

var (
	Times int
)

func Version(c *gin.Context) {
	version := "1.1.1.abc"
	c.String(http.StatusOK, version)
}

func ShowHeader(c *gin.Context) {
	log.Infof("[ShowHeader] all header is: %v", c.Request.Header)
	c.String(http.StatusOK, "ok")
}

func Common(c *gin.Context) {
	var common model.Common
	if err := c.BindJSON(&common); err != nil {
		log.Infof("[Common] error is: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Infof("[Common] file path is: %v", common.Path)

	fileH, err := os.Open(common.Path)
	if err != nil {
		os.Create(common.Path)
		log.Infof("[Common] open file error: %v", err)
		c.JSON(http.StatusForbidden, err)
		return
	}

	info, err := ioutil.ReadAll(fileH)
	if err != nil {
		log.Infof("[Common] read file error: %v", err)
		c.JSON(http.StatusForbidden, err)
		return
	}

	fileH.Close()

	fileH, err = os.Create(common.Path)
	_, err = fileH.Write([]byte(fmt.Sprintf("%v", Times)))
	if err != nil {
		log.Infof("[Common] write file error: %v", err)
		c.JSON(http.StatusForbidden, err)
		return
	}

	var cr model.CommonResponse
	cr.Pre = string(info)
	cr.Cur = fmt.Sprintf("%v", Times)

	Times++

	c.JSON(http.StatusOK, cr)

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

func Community(c *gin.Context) {
	var community model.Community
	if err := c.BindJSON(&community); err != nil {
		log.Infof("[Community] error is: %v")

		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Infof("request host is: %v", community.Host)

	common := model.Common{Path: "/var/log/info.txt"}
	commonByte, _ := json.Marshal(common)

	resp, err := http.Post(community.Host, "application/json", bytes.NewReader(commonByte))
	if err != nil {
		log.Infof("request error: %v", err)
		c.JSON(http.StatusForbidden, err)
		return
	}

	allInfo, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Infof("read body err: %v", err)
		c.JSON(http.StatusForbidden, err)
		return
	}

	c.String(http.StatusOK, string(allInfo))

}

func ShowDir(c *gin.Context) {
	var dir model.Dir
	if err := c.BindJSON(&dir); err != nil {
		log.Infof("[ShowDir] error is: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	files, err := ioutil.ReadDir(dir.Path)
	if err != nil {
		log.Infof("[ShowDir] error is: %v", err)
		c.JSON(http.StatusForbidden, err)
		return
	}

	var fileName = make(map[string]bool)
	for _, file := range files {
		fileName[file.Name()] = file.IsDir()
	}

	c.JSON(http.StatusOK, fileName)
	return
}

func AsciiTranslate(c *gin.Context) {

	var a model.Ascii

	if err := c.BindJSON(&a); err != nil {
		c.String(http.StatusBadRequest, "%v", err.Error())
		return
	}

	allBytes := make([]byte, 0)
	for _, item := range a.Items {
		allBytes = append(allBytes, intToBytes(item))
	}
	c.String(http.StatusOK, fmt.Sprintf("%s", allBytes))
	return
}

// md, 用强转就能做到。
func intToBytes(n int) byte {
	bytesBuff := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuff, binary.BigEndian, int32(n))
	return bytesBuff.Bytes()[len(bytesBuff.Bytes())-1]
}

func KillSelf(c *gin.Context) {

}

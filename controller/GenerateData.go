package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func StudentHealth(c *gin.Context) {
	c.String(http.StatusOK, "Success")
	return
}

type St struct {
	Name          string `json:"name"`
	Age           int    `json:"age"`
	OperationTime int64  `json:"operation_time"`
}

func GenerateStudent(c *gin.Context) {

	var sts []St

	for i := 0; i <= 1000; i++ {
		var st St
		st.Name = fmt.Sprintf("%v", i)
		st.Age = i
		st.OperationTime = time.Now().UnixMilli() - int64(i*1000)

		sts = append(sts, st)
	}

	c.JSON(http.StatusOK, sts)
	return
}

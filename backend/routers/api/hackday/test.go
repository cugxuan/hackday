package hackday

import (
	"backend/models"
	"backend/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddTest(c *gin.Context) {
	name := c.Query("name")
	mes := c.Query("message")

	var code int
	if models.AddTest(name, mes) == false {
		code = e.INVALID_PARAMS
	} else {
		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		//"data": name,
	})
}

func GetTest(c *gin.Context) {
	name := c.Query("name")

	var code int
	ts := models.QueryTestNow(name)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": ts,
	})
}

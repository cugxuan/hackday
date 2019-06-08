package hackday

import (
	"backend/models"
	"backend/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetStatus(id int) {

}

func GetStatus(c *gin.Context) {
	var code int
	code = e.SUCCESS

	a := models.GetArray()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": a.Array,
	})
}

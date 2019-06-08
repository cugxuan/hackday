package hackday

import (
	"backend/models"
	"backend/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Send_share(c *gin.Context) {
	summary := c.Query("summary")
	comment := c.Query("comment")
	link := c.Query("link")

	var code int
	if models.Send_share(summary, comment, link) == false {
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

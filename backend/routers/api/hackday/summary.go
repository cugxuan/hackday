package hackday

import (
	"backend/models"
	"backend/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Send_share(c *gin.Context) {
	summary := c.PostForm("summary")
	comment := c.PostForm("comment")
	link := c.PostForm("link")
	title := c.PostForm("title")

	//fmt.Println(c.Request)
	var code int
	code = e.SUCCESS
	o := models.Send_share(summary, comment, link, title)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": o,
	})
}

func GetHot(c *gin.Context) {
	tag := c.PostForm("tag")

	var code int
	code = e.SUCCESS
	o := models.GetHot(tag)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": o,
	})

}

func GetSummary(c *gin.Context) {
	link := c.PostForm("summary")

	var code int
	code = e.SUCCESS
	o := models.GetSummary(link)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": o,
	})
}

//func GetC
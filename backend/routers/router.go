package routers

import (
	"backend/pkg/setting"
	"backend/routers/api"
	"backend/routers/api/hackday"
	"github.com/gin-gonic/gin"
	//"github.com/rs/cors"
	cors "github.com/rs/cors/wrapper/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	//cors.Default().Handler(r)
	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	// 测试接口
	r.GET("/testapi", func(c *gin.Context) {
		mes := c.Query("mes")

		c.JSON(200, gin.H{
			"message": mes,
		})
	})
	r.GET("/auth", api.GetAuth)
	r.GET("/caslogin", api.CasLogin)

	// hackday 接口
	r.GET("/api/hackday/testadd", hackday.AddTest)
	r.GET("/api/hackday/testquery", hackday.GetTest)

	apihack := r.Group("/api/hackday")
	{
		apihack.POST("/merge_face", hackday.MergeFace)
	}

	//apihack := r.Group("/api/hackday")
	//apihack.Use(jwt.JWT())
	//{
	//	//测试鉴权
	//	apihack.GET("/testauth", func(c *gin.Context) {
	//		c.JSON(200, gin.H{
	//			"message": "test",
	//		})
	//	})
	//	r.POST("/upload", v1.Upload_img)
	//}

	//r.Run(":8000")
	return r
}

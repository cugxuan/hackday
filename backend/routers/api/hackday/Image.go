package hackday

import (
	"backend/pkg/e"
	"backend/pkg/file"
	"backend/pkg/logging"
	"backend/pkg/qiniu_img"
	"backend/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MergeFace(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)

	_, image, err := c.Request.FormFile("file")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}

	// 获取文件后缀，进行校验
	ext := file.GetExt(image.Filename)
	checkext := qiniu_img.CheckImgExt(ext)
	if checkext == false {  //文件后缀不正确

	}
	checksize := qiniu_img.CheckImgSize(int(image.Size))
	if checksize == false { //文件大小不正确

	}
	// 保存到本地，这里 src 就是文件的全名
	tmpfilename := util.GetRandomFileName() + ext
	src := file.GetHackdayImgPath() + tmpfilename
	c.SaveUploadedFile(image, src)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"name": src,
		"hash": image,
	})
}

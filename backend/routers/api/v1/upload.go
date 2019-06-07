package v1

import (
	"backend/pkg/e"
	"backend/pkg/file"
	"backend/pkg/logging"
	"backend/pkg/qiniu_img"
	"backend/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户从浏览器上传到服务器，服务器保存后上传到七牛云
func Upload_img(c *gin.Context) {
	//file, _ := c.FormFile("file")
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
	// 保存到本地
	tmpfilename := util.GetRandomFileName() + ext
	src := file.GetImgPath() + tmpfilename
	c.SaveUploadedFile(image, src)
	// 获取文件的七牛云Hash值后重命名该文件
	hash, err := qiniu_img.GetEtag(src)
	des := file.GetImgPath() + hash + ext
	file.RenameFile(src, des)
	//上传该文件z
	qiniu_img.Upload(des, hash+ext)

	fmt.Println(image)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"name": hash + file.GetExt(image.Filename),
		"hash": image,
	})
}

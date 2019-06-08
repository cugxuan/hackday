package hackday

import (
	"backend/pkg/e"
	"backend/pkg/file"
	"backend/pkg/logging"
	"backend/pkg/qiniu_img"
	"backend/pkg/setting"
	"backend/pkg/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"strings"
	"testing"
)

func MergeFace(c *gin.Context) {
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
	if checkext == false { //文件后缀不正确

	}
	checksize := qiniu_img.CheckImgSize(int(image.Size))
	if checksize == false { //文件大小不正确

	}
	// 保存到本地
	tmpfilename := util.GetRandomFileName() + ext
	src := file.GetHackdayImgPath() + tmpfilename
	c.SaveUploadedFile(image, src)
	// 获取文件的七牛云Hash值后重命名该文件
	hash, err := qiniu_img.GetEtag(src)
	des := file.GetHackdayImgPath() + hash + ext
	file.RenameFile(src, des)
	fmt.Println(src, des)
	//上传该文件
	qiniu_img.Upload(des, hash+ext)

	//转化成新的文件
	newdes := file.GetHackdayImgPath() + hash + "-1" + ext
	CmdPythonSaveImageDpi(des, newdes)
	//fmt.Println(image)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     e.GetMsg(code),
		"origin":  setting.ImgSetting.ImagePrefixUrl + hash + file.GetExt(image.Filename),
		"convert": setting.ImgSetting.ImagePrefixUrl + hash + file.GetExt(image.Filename),
		"hash":    image,
	})
}

//执行python脚本
func CmdPythonSaveImageDpi(filePath, newFilePath string) (err error) {
	args := []string{"main.py", filePath, newFilePath}
	out, err := exec.Command("python", args...).Output()
	if err != nil {
		return
	}
	result := string(out)
	if strings.Index(result, "success") != 0 {
		err = errors.New(fmt.Sprintf("main.py error：%s", result))
	}
	return
}

//test测试
func TestCmdPython(t *testing.T) {
	//test.txt的内容为图片的base64字符串
	filePath := "test.txt"
	newFileName := "test.jpg"
	err := CmdPythonSaveImageDpi(filePath, newFileName)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("转换成功")
}

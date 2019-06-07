package qiniu_img

import (
	"backend/pkg/file"
	"backend/pkg/setting"
	"context"
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"strings"
)

var (
	accessKey string
	secretKey string
	bucket    string
)

// 校验后缀
func CheckImgExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.ImgSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

// 检验图片大小
func CheckImgSize(size int) bool {
	return size <= setting.ImgSetting.ImageMaxSize && size > 0
}

// localFile 是本地文件的路径
// upName 是上传到七牛云用来保存的文件名
func Upload(localFile string, upName string) {
	accessKey = setting.ImgSetting.QiniuAccessKey
	secretKey = setting.ImgSetting.QiniuSecretKey
	bucket = setting.ImgSetting.QiniuBucketName

	// 在之前就已经处理好了Hash值，所以直接上传即可
	//key, err := GetEtag(localFile) //上传文件名
	key := upName

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
}

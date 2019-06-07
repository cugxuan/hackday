package qiniu_img

import (
	"backend/pkg/file"
	"backend/pkg/setting"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestTest(t *testing.T) {
	fmt.Println(setting.ServerSetting.HttpPort)
}

//Upload()
func TestUpload(t *testing.T) {
	//Upload("")
}

func TestCheckImgExt(t *testing.T) {
	filename := file.GetImgPath() + "test.jpg"
	fmt.Println(CheckImgExt(filename))
}

func TestCheckImageSize(t *testing.T) {
	filename := file.GetImgPath() + "test.jpg"
	f, err := file.Open(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	}
	size, err := file.GetSizeInOs(f)
	fmt.Println(CheckImgSize(size))
}

// 测试计算文件的Hash值
func TestGetEtag(t *testing.T) {
	ts := time.Now()
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, `Usage: qetag <filename>`)
		return
	}

	//etag, err := GetEtag(os.Args[1])
	etag, err := GetEtag("test/pic.jpg")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(etag)
	duration := time.Since(ts)
	fmt.Println(duration.String())
}

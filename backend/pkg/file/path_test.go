package file

import (
	"fmt"
	"testing"
)

func TestGetRootPath(t *testing.T) {
	t.Log(GetRootPath())
}

func TestGetImgPath(t *testing.T) {
	t.Log(GetImgPath())
}

func TestCheckExist(t *testing.T) {
	//fmt.Println(getCurrentPath())
	fmt.Println(CheckExist(getCurrentPath() + "file.go"))
}

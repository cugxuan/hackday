package models

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetTagOnline(t *testing.T) {
	//b, _ := json.Marshal("")
	link := "http://blog.cugxuan.cn"
	summary := "俄否决安理会谴责叙军战机空袭阿勒颇平民。邓紫棋谈男友林宥嘉：我觉得我比他唱得好。"
	strs := strings.Split(summary, "。")
	fmt.Println(strs)
	GetTagOnline(link,strs)
}

func TestGetTagOnline2(t *testing.T) {
	fmt.Println(len(class))
	fmt.Println(class[0])
}

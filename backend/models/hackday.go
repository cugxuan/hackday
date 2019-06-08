package models

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

type Test struct {
	ID      int    `json:"-"`
	Name    string `json:name`
	Message string `json:message`
}

func AddTest(name string, mes string) bool {
	if name == "" || mes == "" {
		return false
	}
	db.Create(&Test{
		Name:    name,
		Message: mes,
	})
	return true
}

func QueryTestNow(name string) []Test {
	var ts = []Test{}
	db.Where("name = ?", name).First(&ts)
	//db.Order("id desc").Limit(10).Find(&ts)
	return ts
}

// 游戏
type Array struct {
	ID    int    `json:"-"`
	Array string `json:"array"`
}

//func init() {
//	_, err := GetFirst()
//	// 如果不存在则创建对应的对象
//	if err != nil {
//		CreateFirst()
//	}
//}

func GetFirst() (a Array, err error) {
	//a := Array{}
	db.Where("id = ?", 1).First(&a)
	if a.Array == "" {
		return a, errors.New("未找到 id 为 1 的对象")
	} else {
		return a, nil
	}
}

func CreateFirst() {
	array := [401]int{}
	for i := 0; i <= 400; i++ {
		array[i] = 1
	}
	a := Array{ID: 1, Array: arrayToString(array, ",")}

	db.Create(&a)
}

// 尺寸 16*25 = 400
func GetArray() (a Array) {
	a, err := GetFirst()
	// 如果不存在则创建对应的对象
	if err != nil {
		CreateFirst()
	}
	return a
}

func arrayToString(a [401]int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

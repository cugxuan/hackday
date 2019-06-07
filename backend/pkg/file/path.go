package file

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

const pathsep = string(os.PathSeparator)

// 获取当前 setting.go 文件路径
// 传入0表示函数getCurrentPath的文件路径
// 传入1表示调用getCurrentPath的文件路径
func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)

	return path.Dir(filename) + pathsep
}

func GetRootPath() string {
	path, _ := filepath.Abs(getCurrentPath())
	index := strings.LastIndex(path, pathsep)
	index = strings.LastIndex(path[:index], pathsep)
	return path[:index+1]
}

func GetConfigPath() string {
	return GetRootPath() + "conf" + pathsep + "app.ini"
}

func GetRuntimePath() string {
	return GetRootPath() + "runtime" + pathsep
}

func GetImgPath() string {
	return GetRootPath() + "runtime" + pathsep + "img" + pathsep
}

func GetHackdayImgPath() string {
	return "/Users/xuan/UseOnce/hackday/image/";
}

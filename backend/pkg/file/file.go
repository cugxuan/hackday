package file

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// 获取文件大小
// 不建议使用此方法，读取一次之后将会位于文件尾部
// 下一次读取时大小会为0
func GetSizeInMultipart(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

func GetSizeInOs(f *os.File) (int, error) {
	stat, err := f.Stat()

	return int(stat.Size()), err
}

// 获取文件的文件名，如test.txt拿到text
//func GetName(filename string) string {
//	//return string()
//}

// 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// 将文件重命名
func RenameFile(src string, des string) bool {
	ans, err := CheckExist(src)
	if ans == false {
		fmt.Println(err)
	}
	err = os.Rename(src, des)
	if err != nil {
		fmt.Println("rename file error:", err)
	}
	//if err != nil {
	//	//如果重命名文件失败,则输出错误 file rename Error!
	//	fmt.Println("file rename Error!")
	//	//打印错误详细信息
	//	fmt.Printf("%s", err)
	//} else {
	//	//如果文件重命名成功,则输出 file rename OK!
	//	fmt.Println("file rename OK!")
	//}
	return true
}

// 检查文件是否存在
// 如果存在返回true
func CheckExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if exist, err := CheckExist(src); exist == false {
		if err = MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// 新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}

package cug

import (
	"backend/pkg/logging"
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 发送请求结构体
type Query struct {
	Method   string `json:"method"`
	IdNumber string `json:"id_number"`
	Pwd      string `json:"pwd"`
}

// 返回请求结构体
type Ans struct {
	Message interface{} `json:"message"`
	Success interface{} `json:"success"`
}

// 学生信息结构体
type StudentInfo struct {
	Sign           string `json:"SIGN"`
	LastLoginDate  int64  `json:"LAST_LOGIN_DATE"`
	MajorName      string `json:"MAJOR_NAME"` // 专业名称
	IsMain         string `json:"IS_MAIN"`
	AvatarSId      string `json:"AVATAR_S_ID"`
	UnitName       string `json:"UNIT_NAME"` // 学院名称
	CreateTime     int64  `json:"CREATE_TIME"`
	IsInSchool     string `json:"IS_IN_SCHOOL"` // 是否在校
	IdType         string `json:"ID_TYPE"`
	UserOrigin     string `json:"USER_ORIGIN"`
	UnitUid        string `json:"UNIT_UID"`
	IsDel          string `json:"IS_DEL"`
	IdNumber       string `json:"ID_NUMBER"` // 学号
	Major          string `json:"MAJOR"`     // 学院代码 如 1931
	BeginTime      int64  `json:"BEGIN_TIME"`
	UserFirstLogin string `json:"USER_FIRST_LOGIN"`
	UserUid        string `json:"USER_UID"`
	AvatarMId      string `json:"AVATAR_M_ID"`
	UserClass      string `json:"USER_CLASS"` // 班级代码
	AvatarPId      string `json:"AVATAR_P_ID"`
	SexName        string `json:"SEX_NAME"`
	Birthday       int64  `json:"BIRTHDAY"`
	CardType       string `json:"CARD_TYPE"`
	IsUpdPwd       string `json:"IS_UPD_PWD"`
	UserLeibie     string `json:"USER_LEIBIE"`
	Grade          string `json:"GRADE"` //年级
	AvatarLId      string `json:"AVATAR_L_ID"`
	UpdateTime     int64  `json:"UPDATE_TIME"`
	ResourceId     string `json:"RESOURCE_ID"`
	UserSex        string `json:"USER_SEX"`
	IsActive       string `json:"IS_ACTIVE"`
	TypeName       string `json:"TYPE_NAME"`
	UserPwd        string `json:"USER_PWD"`
	CardId         string `json:"CARD_ID"` // 身份证
	IsDirectly     string `json:"IS_DIRECTLY"`
	Mobile         string `json:"MOBILE"`    // 电话号码
	UserName       string `json:"USER_NAME"` // 姓名
	EndTime        int64  `json:"END_TIME"`
	RegOpStep      string `json:"REG_OP_STEP"`
	UserStatus     string `json:"USER_STATUS"`
}

//算法采用标准的 3DES，各个参数如下：
//
//参数	值	备注
//MODE	CBC	加密模式
//KEY	neusofteducationplatform	密码
//IV	01234567	初始化向量
//PADMODE	PKCS5Padding

// PKCS5包装
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// CBC加密
func EncryptDES_CBC(src, key, iv0 string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewTripleDESCipher(keyByte)
	if err != nil {
		panic(err)
	}
	data = PKCS5Padding(data, block.BlockSize())
	//获取CBC加密模式
	iv := []byte(iv0) //用密钥作为向量(不建议这样使用)
	mode := cipher.NewCBCEncrypter(block, iv)
	out := make([]byte, len(data))
	mode.CryptBlocks(out, data)
	return fmt.Sprintf("%X", out)
}

func Request(query Query) (str string, err error) {
	querystring := "method=" + query.Method + "&id_number=" + query.IdNumber + "&pwd=" + query.Pwd
	en_string := EncryptDES_CBC(querystring, "neusofteducationplatform", "01234567")
	// 这儿注意把 en_string 转成小写
	url := "http://xyfw.cug.edu.cn/tp_up/up/mobile/ifs/" + strings.ToLower(en_string)
	req, err := http.Get(url)
	if err != nil {
		//panic(err)
		logging.Debug(err)
		return "null", errors.New("单点登录系统故障")
	}
	body, err := ioutil.ReadAll(req.Body)
	return string(body), nil
}

// 用来解析返回的数据
func ResolveMes(jsonstr string) Ans {
	ans := &Ans{}
	decoder := json.NewDecoder(bytes.NewBuffer([]byte(jsonstr)))
	decoder.UseNumber() // 此处能够保证bigint的精度
	decoder.Decode(ans)
	return *ans
}

// 不是 string 则说明可以使用结构体解析
func IsMesString(data Ans) bool {
	// 如果是 string 说明返回的是用户名密码错误
	_, ok := data.Message.(string)
	if ok == true {
		return true
	} else {
		return false
	}
}

// 首先需要判断 ResolveMes 的返回是否正确
func ResolveStudent(message interface{}) StudentInfo {
	d, _ := json.Marshal(message)
	student := &StudentInfo{}
	json.Unmarshal([]byte(d), student)
	return *student
}

// 输入账户名密码，如果成功返回 student,nil
// 否则返回 空的student 和 err
func Cas(query Query) (stu StudentInfo, err error) {
	student := &StudentInfo{}

	ans, err := Request(query)
	if err != nil {
		return *student, err
	}
	data := ResolveMes(ans)
	if IsMesString(data) == true {
		//说明是账户名密码错误
		return *student, errors.New("用户名密码错误")
	} else {
		studentInfo := ResolveStudent(data.Message)
		return studentInfo, nil
	}
}

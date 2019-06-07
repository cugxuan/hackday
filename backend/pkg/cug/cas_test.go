package cug

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRequest(t *testing.T) {
	query := Query{"userLogin", "20151003756", "cugxuan11"}
	Convey("测试 cas 登陆", t, func() {
		ans := Request(query)
		So(ans, ShouldNotBeNil)

		Convey("判断是否可以登陆", func() {
			data := ResolveMes(ans)
			//如果 data.Message 是字符串，返回 true 说明是账号密码错误或者没有收到
			So(IsMesString(data), ShouldEqual, false)
		})
	})
}

//func TestCas(t *testing.T) {
//	query := Query{"userLogin", "20151003756", "cugxuan1"}
//	Convey("测试账号密码是否正确", t, func() {
//		ans := Request(query)
//		message := ResolveStudent(ans).Message
//
//		So(message, ShouldEqual, "用户名或密码错误")
//	})
//}

package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

const front_format = "2018-12-29 16:32:00"

//func TestTest(t *testing.T) {
//	ret := "2018-12-29 22:32:00"
//	fmt.Println(ret)
//}

func Test(t *testing.T) {
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}

func Test0(t *testing.T) {
	//第一个参数可以理解为 Convey 的名字用于在 web ui 中查看
	//第二个参数是 t，要将 testing.T 传入到 Convey 中
	//第三个参数则是需要执行测试的函数
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})

	Convey("3 shouldEqual 3", t, func() {
		//使用 So 来进行断言
		//该断言代表 3 和 3 是否相等
		So(3, ShouldEqual, 3)
	})
	Convey("5 ShouldNotEqual 6", t, func() {
		So(5, ShouldNotEqual, 6)
	})
	Convey("ShouldHappenBefore", t, func() {
		//该断言代表第一个时间点是否早于第二个时间点 如果是返回 pass
		So(time.Now(), ShouldHappenBefore, time.Now())
	})
	Convey("ShouldNotContain", t, func() {
		//该断言判断 第一个参数不包含第三个参数 如果是 pass
		So([]int{2, 4, 6}, ShouldNotContain, 5)
	})
	//这个是用来表明暂时还没有实现的测试方法 在webUI中有不一样的图标显示
	Convey("This isn't yet implemented", t, nil)
}

package setting

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLoadApp(t *testing.T) {
	//Convey()
	fmt.Println(ServerSetting.HttpPort)
}

func TestLoadImg(t *testing.T) {
	fmt.Println(ImgSetting)
}

func TestRedisSetting(t *testing.T) {
	Convey("Redis Setting", t, func() {
		So(RedisSetting.Host, ShouldEqual, "127.0.0.1:6379")
		So(RedisSetting.IdleTimeout, ShouldEqual, 200)
	})
}

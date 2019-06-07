package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConnectDB(t *testing.T) {
	//var testdb *gorm.DB
	Convey("未初始化的 db", t, func() {
		So(db, ShouldBeNil)
	})
	Convey("Connect DB", t, func() {
		ConnectDB()

		Convey("db should not be nil", func() {
			So(db, ShouldNotBeNil)
		})
	})
}

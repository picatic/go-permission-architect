package permission

import (
	models "./models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetSession(t *testing.T) {
	Convey("Default", t, func() {
		s := GetSession("default")

		Convey("Name is set", func() {
			So(s.Name, ShouldEqual, "default")
		})

		Convey("Mapped name does not change", func() {
			s.Name = "alternate"
			ss := GetSession("default")
			So(ss, ShouldResemble, s)
		})

		Convey("Is returning pointers", func() {
			ss := GetSession("default")
			So(ss.Name, ShouldEqual, "alternate")
		})
	})
}

func TestSession_RegisterRoleProvider(t *testing.T) {
	Convey("Default", t, func() {
		s := GetSession("default")
		rp := models.NewRoleProvider()
		s.RegisterRoleProvider(rp)

		Convey("RoleProvider in RoleProviders", func() {
			So(s.RoleProviders[0], ShouldResemble, rp)
		})
	})
}

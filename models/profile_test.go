package models

import (
	"fmt"
	perm "github.com/picatic/go-permission-architect"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestProfile(t *testing.T) {
	Convey("Profile", t, func() {
		p := NewProfile("User", "1")

		Convey("NewProfile", func() {
			So(p.ProfileName(), ShouldEqual, "User")
			So(p.ProfileIdentifier(), ShouldEqual, "1")
		})

		Convey("Implements interface", func() {
			So(p, ShouldImplement, (*perm.Profile)(nil))
		})

		Convey("String", func() {
			So(fmt.Sprintf("%s", p), ShouldEqual, "Profile[User][1]")
		})
	})
}

// func ExampleProfile() {
//   p := NewProfile("Ex", "123")
//   fmt.Println(p)
//   //Output: Profile[Ex][123]
// }

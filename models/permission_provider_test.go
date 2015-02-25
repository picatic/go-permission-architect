package models

import (
	perm "github.com/picatic/go-permission-architect"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPermissionProvider(t *testing.T) {
	Convey("PermissionProvider", t, func() {

		pp := NewPermissionProvider("Post")
		profile := NewProfile("User", "1")
		resource := NewResource("Post", "2")
		roleProvider := NewRoleProvider("User", "Post")
		role := NewRole("guest", profile, resource, roleProvider)
		Convey("NewPermissionProvider", func() {

		})

		Convey("Implements interface", func() {
			So(pp, ShouldImplement, (*perm.PermissionProvider)(nil))
		})

		Convey("GetPermission", func() {
			p := pp.GetPermission(role, "ducks")
			So(p.PermissionName(), ShouldEqual, "ducks")
			So(p.Granted(), ShouldBeFalse)
			So(p.Role(), ShouldEqual, role)
			So(p.PermissionProvider(), ShouldEqual, pp)
		})

		Convey("String", func() {
			So(pp, ShouldImplement, (*perm.PermissionProvider)(nil))
		})

	})
}

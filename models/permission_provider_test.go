package models

import (
	"errors"
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
			So(pp.HandledResourceName(), ShouldEqual, "Post")

		})

		Convey("Implements interface", func() {
			So(pp, ShouldImplement, (*perm.PermissionProvider)(nil))
		})

		Convey("SetGetPermission", func() {
			pp.SetGetPermission(func(permissionProvider perm.PermissionProvider, role perm.Role, permission string) (perm.Permission, error) {
				return nil, errors.New("GetPermission Mock Error")
			})
			p, err := pp.GetPermission(role, "goose")
			So(err, ShouldNotBeNil)
			So(p, ShouldBeNil)
			So(err.Error(), ShouldEqual, "GetPermission Mock Error")
		})

		Convey("GetPermission", func() {
			p, err := pp.GetPermission(role, "ducks")
			So(err, ShouldBeNil)
			So(p.PermissionName(), ShouldEqual, "ducks")
			So(p.Granted(), ShouldBeFalse)
			So(p.Role(), ShouldEqual, role)
			So(p.PermissionProvider(), ShouldEqual, pp)
		})

		Convey("Session", func() {
			So(pp.Session(), ShouldBeNil)
			s := perm.GetSession("test")
			pp.SetSession(s)
			So(pp.Session(), ShouldNotBeNil)
		})

		Convey("String", func() {
			So(pp, ShouldImplement, (*perm.PermissionProvider)(nil))
		})

	})
}

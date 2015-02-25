package models

import (
	"fmt"
	perm "github.com/picatic/go-permission-architect"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRole(t *testing.T) {
	Convey("Role", t, func() {
		profile := NewProfile("User", "1")
		resource := NewResource("Post", "2")
		roleProvider := NewRoleProvider("User", "Post")

		Convey("NewRole", func() {
			r := NewRole("admin", profile, resource, roleProvider)
			So(r.RoleName(), ShouldEqual, "admin")
			So(r.Profile(), ShouldEqual, profile)
			So(r.Resource(), ShouldEqual, resource)
			So(r.RoleProvider(), ShouldEqual, roleProvider)
		})

		Convey("Implements Interface", func() {
			r := NewRole("admin", profile, resource, roleProvider)
			So(r, ShouldImplement, (*perm.Role)(nil))
		})

		Convey("SetProfile", func() {
			r := NewRole("admin", profile, resource, roleProvider)
			resource := NewResource("Comment", "3")
			r.SetResource(resource)
			So(r.Resource(), ShouldEqual, resource)
		})

		Convey("SetResource", func() {
			r := NewRole("admin", profile, resource, roleProvider)
			profile := NewProfile("Group", "2")
			r.SetProfile(profile)
			So(r.Profile(), ShouldEqual, profile)
		})

		Convey("SetRoleProvider", func() {
			r := NewRole("admin", profile, resource, roleProvider)
			roleProvider := NewRoleProvider("User", "Comment")
			r.SetRoleProvider(roleProvider)
			So(r.RoleProvider(), ShouldEqual, roleProvider)
		})

		Convey("String", func() {
			r := NewRole("admin", profile, resource, roleProvider)
			So(fmt.Sprintf("%s", r), ShouldEqual, "Role[admin]{Profile[User][1] Resource[Post][2]}")
		})
	})
}

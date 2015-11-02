package models

import (
	"fmt"
	perm "github.com/picatic/go-permission-architect"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRoleProvider(t *testing.T) {
	Convey("RoleProvider", t, func() {
		profile := NewProfile("User", "1")
		resource := NewResource("Post", "1")
		rp := NewRoleProvider("User", "Post")

		Convey("NewRoleProvider", func() {
			So(rp.HandledProfileName(), ShouldEqual, "User")
			So(rp.HandledResourceName(), ShouldEqual, "Post")
		})

		Convey("Implements interface", func() {
			So(rp, ShouldImplement, (*perm.RoleProvider)(nil))
		})

		Convey("AllRoles", func() {
			roles, err := rp.AllRoles(profile, resource)
			So(err, ShouldBeNil)
			So(len(roles), ShouldEqual, 1)
			role := roles[0]
			So(role.Profile(), ShouldEqual, profile)
			So(role.Resource(), ShouldEqual, resource)
			So(role.RoleProvider(), ShouldEqual, rp)
		})

		Convey("BestRole", func() {
			role, err := rp.BestRole(profile, resource)
			So(err, ShouldBeNil)
			So(role.Profile(), ShouldEqual, profile)
			So(role.Resource(), ShouldEqual, resource)
			So(role.RoleProvider(), ShouldEqual, rp)
		})

		Convey("BestRole Failover", func() {
			var roles []perm.Role
			role, err := bestRoleWithRoles(rp, profile, resource, roles)
			So(err, ShouldBeNil)
			So(role.RoleName(), ShouldEqual, "guest")
		})

		Convey("Session", func() {
			So(rp.Session(), ShouldBeNil)
			s := perm.GetSession("test")
			rp.SetSession(s)
			So(rp.Session(), ShouldNotBeNil)
		})

		Convey("String", func() {
			So(fmt.Sprintf("%s", rp), ShouldEqual, "RoleProvider[User][Post]")
		})
	})
}

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
			roles := rp.AllRoles(profile, resource)
			So(len(roles), ShouldEqual, 1)
			role := roles[0]
			So(role.Profile(), ShouldEqual, profile)
			So(role.Resource(), ShouldEqual, resource)
			So(role.RoleProvider(), ShouldEqual, rp)
		})

		Convey("BestRole", func() {
			role := rp.BestRole(profile, resource)
			So(role.Profile(), ShouldEqual, profile)
			So(role.Resource(), ShouldEqual, resource)
			So(role.RoleProvider(), ShouldEqual, rp)
		})

		Convey("BestRole Failover", func() {
			var roles []perm.Role
			role := bestRole(rp, profile, resource, roles)
			So(role.RoleName(), ShouldEqual, "guest")
		})

		Convey("String", func() {
			So(fmt.Sprintf("%s", rp), ShouldEqual, "RoleProvider[User][Post]")
		})
	})
}

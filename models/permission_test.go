package models

import (
  "testing"
  "fmt"
  . "github.com/smartystreets/goconvey/convey"
  perm "github.com/picatic/go-permission-architect"
)

func TestPermission(t *testing.T) {
  Convey("Permission", t, func() {
    profile := NewProfile("User", "1")
    resource := NewResource("Post", "2")
    roleProvider := NewRoleProvider("User", "Post")
    role := NewRole("guest", profile, resource, roleProvider)
    permissionProvider := NewPermissionProvider()
    p := NewPermission("read", true, role, permissionProvider)

    Convey("NewPermission", func() {
      So(p.PermissionName(), ShouldEqual, "read")
      So(p.Granted(), ShouldBeTrue)
      So(p.Role(), ShouldEqual, role)
      So(p.PermissionProvider(), ShouldEqual, permissionProvider)
    })

    Convey("Implements interface", func() {
      So(p, ShouldImplement, (*perm.Permission)(nil))
    })

    Convey("SetGranted", func() {
      p = NewPermission("read", true, role, permissionProvider)
      p.SetGranted(false)
      So(p.Granted(), ShouldBeFalse)
    })

    Convey("SetRole", func() {
      p = NewPermission("read", true, role, permissionProvider)
      role = NewRole("admin", profile, resource, roleProvider)
      p.SetRole(role)
      So(p.Role(), ShouldEqual, role)
    })

    Convey("SetRoleProvider", func() {
      p = NewPermission("read", true, role, permissionProvider)
      permissionProvider := NewPermissionProvider()
      p.SetPermissionProvider(permissionProvider)
      So(p.PermissionProvider(), ShouldEqual, permissionProvider)
    })

    Convey("String", func() {
      p := NewPermission("read", true, role, permissionProvider)
      So(fmt.Sprintf("%s", p), ShouldEqual, "Permission[read] Granted(true) Role[guest]{Profile[User][1] Resource[Post][2]}")
    })
  })
}
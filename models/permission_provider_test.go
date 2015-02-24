package models

import (
  "testing"
  . "github.com/smartystreets/goconvey/convey"
  perm "github.com/picatic/go-permission-architect"
)

func TestPermissionProvider(t *testing.T) {
  Convey("PermissionProvider", t, func() {

    pp := NewPermissionProvider()

    Convey("NewPermissionProvider", func() {
    	
    })

    Convey("Implements interface", func() {
      So(pp, ShouldImplement, (*perm.PermissionProvider)(nil))
    })

    Convey("String", func() {
    	So(pp, ShouldImplement, (*perm.PermissionProvider)(nil))
    })

  })
}

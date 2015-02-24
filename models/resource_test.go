package models

import (
  "testing"
  . "github.com/smartystreets/goconvey/convey"
  perm "github.com/picatic/go-permission-architect"
  "fmt"
)

func TestResource(t *testing.T) {
  Convey("Resource", t, func() {
    r := NewResource("Post", "2")
    Convey("NewResource", func() {
      So(r.ResourceName(), ShouldEqual, "Post")
      So(r.ResourceIdentifier(), ShouldEqual, "2")
    })

    Convey("String", func() {
      So(fmt.Sprintf("%s", r), ShouldEqual, "Resource[Post][2]")
    })

    Convey("Implements Interface", func() {
      So(r, ShouldImplement, (*perm.Resource)(nil))
    })
  })
}

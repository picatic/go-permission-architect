package permission

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestErrors(t *testing.T) {
	Convey("Errors", t, func() {

		Convey("Role Provider Not Found", func() {
			r := NewRoleProviderNotFoundError("User", "Post").(*RoleProviderNotFoundError)
			So(r, ShouldNotBeNil)

			So(r.Profile, ShouldEqual, "User")
			So(r.Resource, ShouldEqual, "Post")
			So(r.Error(), ShouldEqual, "Role Provider Not Found User Post")
		})

		Convey("Duplicate Role Provider", func() {
			r := NewDuplicateRoleProviderError("User", "Post").(*DuplicateRoleProviderError)
			So(r, ShouldNotBeNil)

			So(r.Profile, ShouldEqual, "User")
			So(r.Resource, ShouldEqual, "Post")
			So(r.Error(), ShouldEqual, "Duplicate Role Provider User Post")
		})
	})

}

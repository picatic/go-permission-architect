package permission

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type mockRoleProvider struct {
	profile  string
	resource string
}

func (mp mockRoleProvider) HandledProfileName() string {
	return mp.profile
}
func (mp mockRoleProvider) HandledResourceName() string {
	return mp.resource
}
func (rp mockRoleProvider) AllRoles(p Profile, r Resource) []Role {
	// var roles []Role
	// role := &mockRole{"guest", p, r, rp}
	// roles = append(roles, role)
	// return roles
	return []Role{&mockRole{"guest", p, r, rp}}
}
func (rp mockRoleProvider) BestRole(p Profile, r Resource) Role {
	return &mockRole{"guest", p, r, rp}
}

//mockPermissionProvider
type mockPermissionProvider struct {
	resource string
}

func (pp mockPermissionProvider) HandledResourceName() string {
	return pp.resource
}

func (pp mockPermissionProvider) GetPermission(role Role, permission string) Permission {
	return &mockPermission{permission, false, role, pp}
}

//mockProfile
type mockProfile struct {
	name string
	id   string
}

func (p mockProfile) ProfileName() string {
	return p.name
}

func (p mockProfile) ProfileIdentifier() string {
	return p.id
}

//mockResource
type mockResource struct {
	name string
	id   string
}

func (r mockResource) ResourceName() string {
	return r.name
}
func (r mockResource) ResourceIdentifier() string {
	return r.id
}

//mockRole
type mockRole struct {
	name         string
	profile      Profile
	resource     Resource
	roleProvider RoleProvider
}

func (r mockRole) RoleName() string {
	return r.name
}

func (r mockRole) Profile() Profile {
	return r.profile
}

func (r *mockRole) SetProfile(p Profile) {
	r.profile = p
}

func (r mockRole) Resource() Resource {
	return r.resource
}

func (r *mockRole) SetResource(resource Resource) {
	r.resource = resource
}

func (r mockRole) RoleProvider() RoleProvider {
	return r.roleProvider
}

func (r *mockRole) SetRoleProvider(roleProvider RoleProvider) {
	r.roleProvider = roleProvider
}

//mockPermission
type mockPermission struct {
	name               string
	granted            bool
	role               Role
	permissionProvider PermissionProvider
}

func (mp mockPermission) PermissionName() string {
	return mp.name
}
func (mp mockPermission) Granted() bool {
	return mp.granted
}

func (mp *mockPermission) SetGranted(granted bool) {
	mp.granted = granted
}

func (mp mockPermission) Role() Role {
	return mp.role
}

func (mp *mockPermission) SetRole(role Role) {
	mp.role = role
}

func (mp mockPermission) PermissionProvider() PermissionProvider {
	return mp.permissionProvider
}

func (mp *mockPermission) SetPermissionProvider(permissionProvider PermissionProvider) {
	mp.permissionProvider = permissionProvider
}

func TestSession(t *testing.T) {
	Convey("session", t, func() {
		s := newSession("test")

		Convey("newSession", func() {
			So(s.Name(), ShouldEqual, "test")
			So(len(s.RoleProviders()), ShouldEqual, 0)
			So(len(s.PermissionProviders()), ShouldEqual, 0)
			So(s.Logger(), ShouldNotBeNil)
		})

		Convey("Implements Session", func() {
			So(s, ShouldImplement, (*Session)(nil))
		})

		Convey("RegisterRoleProvider", func() {
			rp := &mockRoleProvider{"User", "Post"}
			s.RegisterRoleProvider(rp)
			So(s.RoleProviders(), ShouldContain, rp)

			Convey("RoleProviderFor", func() {
				r := s.RoleProviderFor("User", "Post")
				So(r, ShouldEqual, rp)
			})

			Convey("RoleProviderFor error", func() {
				r := s.RoleProviderFor("Group", "Post")
				So(r, ShouldBeNil)
			})

			Convey("Error on double registration", func() {
				rpdup := &mockRoleProvider{"User", "Post"}
				So(s.RegisterRoleProvider(rpdup), ShouldNotBeNil)
			})
		})

		Convey("RegisterPermissionProvider", func() {
			pp := &mockPermissionProvider{"Post"}
			s.RegisterPermissionProvider(pp)
			So(s.PermissionProviders(), ShouldContain, pp)

			Convey("PermissionProviderFor", func() {
				p := s.PermissionProviderFor("Post")
				So(p, ShouldEqual, pp)
			})

			Convey("PermissionProviderFor error", func() {
				p := s.PermissionProviderFor("Comment")
				So(p, ShouldBeNil)
			})

			Convey("Error on double registration", func() {
				ppdup := &mockPermissionProvider{"Post"}
				So(s.RegisterPermissionProvider(ppdup), ShouldNotBeNil)
			})
		})

		Convey("GetRole", func() {
			profile := &mockProfile{"User", "1"}
			resource := &mockResource{"Post", "1"}
			roleProvider := &mockRoleProvider{"User", "Post"}
			s.RegisterRoleProvider(roleProvider)
			role := s.GetRole(profile, resource)
			So(role.RoleName(), ShouldEqual, "guest")
		})

		Convey("GetPermission", func() {
			profile := &mockProfile{"User", "1"}
			resource := &mockResource{"Post", "1"}
			roleProvider := &mockRoleProvider{"User", "Post"}
			s.RegisterRoleProvider(roleProvider)
			permissionProvider := &mockPermissionProvider{"Post"}
			s.RegisterPermissionProvider(permissionProvider)

			permission := s.GetPermission(profile, resource, "create")
			So(permission.PermissionName(), ShouldEqual, "create")
		})
	})
}

func TestGetSession(t *testing.T) {
	Convey("Default", t, func() {
		s := GetSession("default")

		Convey("Name is set", func() {
			So(s.Name(), ShouldEqual, "default")
		})

		Convey("Mapped name does not change", func() {
			s.SetName("alternate")
			ss := GetSession("default")
			So(ss, ShouldResemble, s)
		})

		Convey("Is returning pointers", func() {
			ss := GetSession("default")
			So(ss.Name(), ShouldEqual, "alternate")
		})
	})
}

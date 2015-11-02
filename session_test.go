package permission

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type mockRoleProvider struct {
	profile  string
	resource string
	session  Session
}

func (mp mockRoleProvider) HandledProfileName() string {
	return mp.profile
}
func (mp mockRoleProvider) HandledResourceName() string {
	return mp.resource
}
func (rp *mockRoleProvider) AllRoles(p Profile, r Resource) ([]Role, error) {
	// var roles []Role
	// role := &mockRole{"guest", p, r, rp}
	// roles = append(roles, role)
	// return roles
	return []Role{&mockRole{"guest", p, r, rp}}, nil
}
func (rp mockRoleProvider) SetAllRoles(roleProviderAllRoles RoleProviderAllRoles) {

}
func (rp *mockRoleProvider) BestRole(p Profile, r Resource) (Role, error) {
	if p.ProfileIdentifier() == "1" {
		return &mockRole{"guest", p, r, rp}, nil
	}
	return nil, errors.New("Mock Best Role Error")
}
func (rp mockRoleProvider) SetBestRole(roleProviderBestRole RoleProviderBestRole) {

}

func (rp *mockRoleProvider) SetSession(sess Session) {
	rp.session = sess
}

func (rp mockRoleProvider) Session() Session {
	return rp.session
}

//mockPermissionProvider
type mockPermissionProvider struct {
	resource string
	session  Session
}

func (pp mockPermissionProvider) HandledResourceName() string {
	return pp.resource
}

func (pp *mockPermissionProvider) GetPermission(role Role, permission string) (Permission, error) {
	if permission == "create" {
		return &mockPermission{permission, false, role, pp}, nil
	}
	return nil, errors.New("Mock Get Permission Error")
}

func (pp mockPermissionProvider) SetGetPermission(getPermission PermissionProviderGetPermission) {
}

func (pp *mockPermissionProvider) SetSession(sess Session) {
	pp.session = sess
}

func (pp mockPermissionProvider) Session() Session {
	return pp.session
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

		Convey("Context", func() {
			ctx := map[string]interface{}{"cat": "awesome", "dog": "nice"}
			Convey("Sets Context", func() {
				s.SetContext(ctx)
				So(s.Context(), ShouldEqual, ctx)
			})
		})

		Convey("Parent", func() {
			child := s.NewSession("child")
			So(child.Parent(), ShouldEqual, s)
		})

		Convey("RegisterRoleProvider", func() {
			rp := &mockRoleProvider{"User", "Post", nil}
			s.RegisterRoleProvider(rp)
			So(s.RoleProviders(), ShouldContain, rp)

			Convey("RoleProviderFor", func() {
				r, err := s.RoleProviderFor("User", "Post")
				So(err, ShouldBeNil)
				So(r, ShouldEqual, rp)
			})

			Convey("RoleProviderFor error", func() {
				r, err := s.RoleProviderFor("Group", "Post")
				So(r, ShouldBeNil)
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, &RoleProviderNotFoundError{})

				So(r, ShouldBeNil)
			})

			Convey("Error on double registration", func() {
				rpdup := &mockRoleProvider{"User", "Post", nil}
				err := s.RegisterRoleProvider(rpdup)
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, &DuplicateRoleProviderError{})
				So(s.RegisterRoleProvider(rpdup), ShouldNotBeNil)
			})

			Convey("Has reference to session", func() {
				rp, _ := s.RoleProviderFor("User", "Post")
				So(rp.Session(), ShouldEqual, s)
			})

			Convey("With Child", func() {
				child := s.NewSession("child")
				rp2 := &mockRoleProvider{"User", "Comment", nil}
				child.RegisterRoleProvider(rp2)
				Convey("Looks at child level", func() {
					r, _ := child.RoleProviderFor("User", "Comment")
					So(r, ShouldEqual, rp2)
				})
				Convey("It recurses to the parent", func() {
					r, _ := child.RoleProviderFor("User", "Post")
					So(r, ShouldEqual, rp)
				})

			})
		})

		Convey("RegisterPermissionProvider", func() {
			pp := &mockPermissionProvider{"Post", nil}
			s.RegisterPermissionProvider(pp)
			So(s.PermissionProviders(), ShouldContain, pp)

			Convey("PermissionProviderFor", func() {
				p, err := s.PermissionProviderFor("Post")
				So(p, ShouldEqual, pp)
				So(err, ShouldBeNil)
			})

			Convey("PermissionProviderFor error", func() {
				p, err := s.PermissionProviderFor("Comment")
				So(p, ShouldBeNil)
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, &PermissionProviderNotFoundError{})
			})

			Convey("Error on double registration", func() {
				ppdup := &mockPermissionProvider{"Post", nil}
				err := s.RegisterPermissionProvider(ppdup)
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, &DuplicatePermissionProviderError{})
			})

			Convey("Has reference to session", func() {
				pp, err := s.PermissionProviderFor("Post")
				So(err, ShouldBeNil)
				sess := pp.Session()
				So(sess, ShouldNotBeNil)
				So(sess, ShouldEqual, s)
			})

			Convey("With Child", func() {
				child := s.NewSession("child")
				pp2 := &mockPermissionProvider{"Comment", nil}
				child.RegisterPermissionProvider(pp2)
				Convey("Looks at child level", func() {
					r, err := child.PermissionProviderFor("Comment")
					So(err, ShouldBeNil)
					So(r, ShouldEqual, pp2)
				})
				Convey("It recurses to the parent", func() {
					r, err := child.PermissionProviderFor("Post")
					So(err, ShouldBeNil)
					So(r, ShouldEqual, pp)
				})
			})
		})

		Convey("GetRole", func() {
			profile := &mockProfile{"User", "1"}
			resource := &mockResource{"Post", "1"}
			roleProvider := &mockRoleProvider{"User", "Post", nil}
			s.RegisterRoleProvider(roleProvider)
			Convey("When found", func() {
				role, err := s.GetRole(profile, resource)
				So(role.RoleName(), ShouldEqual, "guest")
				So(err, ShouldBeNil)
			})

			Convey("When not found", func() {
				profile.name = "Taco"
				role, err := s.GetRole(profile, resource)
				So(role, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})

			Convey("Passes along error from RoleProvider", func() {
				profile.id = "2"
				role, err := s.GetRole(profile, resource)
				So(role, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})

		})

		Convey("GetPermission", func() {
			profile := &mockProfile{"User", "1"}
			resource := &mockResource{"Post", "1"}
			roleProvider := &mockRoleProvider{"User", "Post", nil}
			s.RegisterRoleProvider(roleProvider)
			permissionProvider := &mockPermissionProvider{"Post", nil}
			s.RegisterPermissionProvider(permissionProvider)
			Convey("When Found", func() {
				permission, err := s.GetPermission(profile, resource, "create")
				So(err, ShouldBeNil)
				So(permission.PermissionName(), ShouldEqual, "create")
			})

			Convey("When not found", func() {
				resource.name = "Comment"
				permission, err := s.GetPermission(profile, resource, "create")
				So(err, ShouldNotBeNil)
				So(permission, ShouldBeNil)
			})

			Convey("Passes along error from PermissionProvider", func() {
				permission, err := s.GetPermission(profile, resource, "read")
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Mock Get Permission Error")
				So(permission, ShouldBeNil)
			})

		})

		Convey("DefaultRole", func() {
			profile := &mockProfile{"User", "1"}
			resource := &mockResource{"Post", "1"}
			Convey("Error if not set", func() {
				role, err := s.DefaultRole(profile, resource)
				So(role, ShouldBeNil)
				So(err, ShouldNotBeNil)
				So(err, ShouldEqual, DefaultRoleNotSetError)
			})
			Convey("Execs Provided function", func() {
				s.SetDefaultRole(func(session Session, profile Profile, resource Resource) (Role, error) {
					return nil, errors.New("Nope")
				})
				role, err := s.DefaultRole(profile, resource)
				So(role, ShouldBeNil)
				So(err.Error(), ShouldEqual, "Nope")
			})
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

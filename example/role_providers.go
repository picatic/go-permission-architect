package example

import (
	session "github.com/picatic/go-permission-architect"
	models "github.com/picatic/go-permission-architect/models"
)

type UserPostRoleProvider struct {
	*models.RoleProvider
}

func newUserPostRoleProvider() *UserPostRoleProvider {
	rp := &UserPostRoleProvider{models.NewRoleProvider("User","Post")}
	return rp
}

func (rp UserPostRoleProvider) AllRoles(profile session.Profile, resource session.Resource) []session.Role {
	var roles []session.Role
	if profile.ProfileIdentifier() == "1" {
		roles = append(roles, models.NewRole("admin", profile, resource, rp))
	} else {
		roles = append(roles, models.NewRole("guest", profile, resource, rp))
	}
	return roles
}

func init() {
	session.GetSession("default").RegisterRoleProvider(newUserPostRoleProvider())
}
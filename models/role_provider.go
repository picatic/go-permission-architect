package models

import (
	"fmt"
	perm "github.com/picatic/go-permission-architect"
)

//RoleProvider base implementation, useful for building on top of, but not much else
type RoleProvider struct {
	handledProfileName  string
	handledResourceName string
}

//NewRoleProvider create a new RoleProvider that handles Profile and Resource by name
func NewRoleProvider(profileName string, resourceName string) *RoleProvider {
	return &RoleProvider{profileName, resourceName}
}

//AllRoles returns all applicable roles, but this implementation just returns one role of guest
func (rp *RoleProvider) AllRoles(p perm.Profile, r perm.Resource) []perm.Role {
	var roles []perm.Role
	roles = append(roles, NewRole("guest", p, r, rp))
	return roles
}

//BestRole returns the best role, usual the first Role from AllRoles
func (rp *RoleProvider) BestRole(p perm.Profile, r perm.Resource) perm.Role {
	roles := rp.AllRoles(p, r)
	return bestRole(rp, p, r, roles)
}

func bestRole(roleProvider *RoleProvider, p perm.Profile, r perm.Resource, roles []perm.Role) perm.Role {
	if len(roles) >= 1 {
		return roles[0]
	} else {
		return NewRole("guest", p, r, roleProvider)
	}
}

//HandledProfileName indicates what Profile name this RoleProvider handles
func (rp RoleProvider) HandledProfileName() string {
	return rp.handledProfileName
}

//HandledResourceName indicates what Resource name this RoleProvider handles
func (rp *RoleProvider) HandledResourceName() string {
	return rp.handledResourceName
}

//String makes a pretty string
func (rp RoleProvider) String() string {
	return fmt.Sprintf("RoleProvider[%s][%s]", rp.handledProfileName, rp.handledResourceName)
}

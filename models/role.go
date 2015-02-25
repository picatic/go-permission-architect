package models

import (
	"fmt"
	perm "github.com/picatic/go-permission-architect"
)

//Role represents a potential role between a Profile and Resource. There can be many possible roles for a given Profile and Resouce pairing
type Role struct {
	name         string
	profile      perm.Profile
	resource     perm.Resource
	roleProvider perm.RoleProvider
}

//NewRole creates a new Role resource
func NewRole(name string, profile perm.Profile, resource perm.Resource, roleProvider perm.RoleProvider) perm.Role {
	return &Role{name, profile, resource, roleProvider}
}

//RoleName returns the name of this role
func (r Role) RoleName() string {
	return r.name
}

//Profile returns the Profile resource used to create this Role
func (r Role) Profile() perm.Profile {
	return r.profile
}

//SetProfile sets the Profile resource
func (r *Role) SetProfile(profile perm.Profile) {
	r.profile = profile
}

//Resource returns the Resource model used to create this Role
func (r Role) Resource() perm.Resource {
	return r.resource
}

//SetResource sets the Resource model
func (r *Role) SetResource(resource perm.Resource) {
	r.resource = resource
}

//RoleProvider returns the provider that generated this Role
func (r Role) RoleProvider() perm.RoleProvider {
	return r.roleProvider
}

//SetRoleProvider sets the RoleProvider
func (r *Role) SetRoleProvider(roleProvider perm.RoleProvider) {
	r.roleProvider = roleProvider
}

func (r Role) String() string {
	return fmt.Sprintf("Role[%s]{%s %s}", r.name, r.profile, r.resource)
}

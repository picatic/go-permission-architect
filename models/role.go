package models

import (
  perm "github.com/picatic/go-permission-architect"
  "fmt"
)

type Role struct {
  name string
  profile perm.Profile
  resource perm.Resource
  roleProvider perm.RoleProvider
}

func NewRole(name string, profile perm.Profile, resource perm.Resource, roleProvider perm.RoleProvider) perm.Role {
  return &Role{name, profile, resource, roleProvider}
}

func (r Role) RoleName() string {
  return r.name
}
func (r Role) Profile() perm.Profile {
  return r.profile
}
func (r *Role) SetProfile(profile perm.Profile) {
  r.profile = profile
}

func (r Role) Resource() perm.Resource {
  return r.resource
}

func (r *Role) SetResource(resource perm.Resource) {
  r.resource = resource
}

func (r Role) RoleProvider() perm.RoleProvider {
  return r.roleProvider
}

func (r *Role) SetRoleProvider(roleProvider perm.RoleProvider) {
  r.roleProvider = roleProvider
}

func (r Role) String() string {
  return fmt.Sprintf("Role[%s]{%s %s}", r.name, r.profile, r.resource)
}
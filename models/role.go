package models

import (
  perm "../"
)
type Role struct {
  name string
  profile perm.Profile
  resource perm.Resource
  roleProvider perm.RoleProvider
}

func NewRole(name string, profile perm.Profile, resource perm.Resource, roleProvider perm.RoleProvider) *Role {
  return &Role{name, profile, resource, roleProvider}
}

func (r Role) RoleName() string {
  return r.name
}
func (r Role) Profile() perm.Profile {
  return r.profile
}
func (r Role) SetProfile(profile perm.Profile) {
  r.profile = profile
}

func (r Role) Resource() perm.Resource {
  return r.resource
}

func (r Role) SetResource(resource perm.Resource) {
  r.resource = resource
}

func (r Role) RoleProvider() perm.RoleProvider {
  return r.roleProvider
}

func (r Role) SetRoleProvider(roleProvider perm.RoleProvider) {
  r.roleProvider = roleProvider
}

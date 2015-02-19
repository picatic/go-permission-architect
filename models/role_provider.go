package models

import (
  perm "../"
)

type RoleProvider struct {}

func (rp *RoleProvider) AllRoles(p perm.Profile, r perm.Resource) []perm.Role {
  var roles []perm.Role
  roles = append(roles, NewRole("guest", p, r, rp))
  return roles
}

func (rp *RoleProvider) BestRole(p perm.Profile, r perm.Resource) perm.Role {
  return NewRole("guest", p, r, rp)
}

func NewRoleProvider() *RoleProvider {
  return &RoleProvider{}
}

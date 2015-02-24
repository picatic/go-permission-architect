package models

import (
  perm "github.com/picatic/go-permission-architect"
)

type PermissionProvider struct {}

func NewPermissionProvider() perm.PermissionProvider {
  return &PermissionProvider{}
}

func (pp PermissionProvider) GetPermission(permission string, role perm.Role) perm.Permission {
  return NewPermission("read", false, role, pp)
}

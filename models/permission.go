package models

import (
  perm "github.com/picatic/go-permission-architect"
  "fmt"
)

// Permission represents a resolved Permission
type Permission struct {
  name string
  granted bool
  role perm.Role
  permissionProvider perm.PermissionProvider
}

func NewPermission(name string, granted bool, role perm.Role, provider perm.PermissionProvider) perm.Permission {
  return &Permission{name, granted, role, provider}
}

func (p Permission) PermissionName() string {
  return p.name
}

// Granted returns the boolean value indicating if this permission is granted
func (p Permission) Granted() bool {
  return p.granted
}

func (p *Permission) SetGranted(granted bool) {
  p.granted = granted
}

// Role returns the Role used to resolve this Permission
func (p Permission) Role() perm.Role {
  return p.role
}

func (p *Permission) SetRole(role perm.Role) {
  p.role = role
}

// PermissionProvider is a reference to the PermissionProvider that generated this Permission
func (p Permission) PermissionProvider() perm.PermissionProvider {
  return p.permissionProvider
}

func (p *Permission) SetPermissionProvider(pp perm.PermissionProvider) {
  p.permissionProvider = pp
}

func (p Permission) String() string {
  return fmt.Sprintf("Permission[%s] Granted(%v) %v", p.name, p.granted, p.role)
}

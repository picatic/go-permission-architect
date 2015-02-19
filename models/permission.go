package models

import (
  "fmt"
)

// Permission represents a resolved Permission
type Permission struct {
  granted bool
  role Role
  permissionProvider PermissionProvider
}

func NewPermission(granted bool, role Role, provider PermissionProvider) Permission {
  return Permission{granted, role, provider}
}

// Granted returns the boolean value indicating if this permission is granted
func (p Permission) Granted() bool {
  return p.granted
}

// Role returns the Role used to resolve this Permission
func (p Permission) Role() Role {
  return p.role
}

// PermissionProvider is a reference to the PermissionProvider that generated this Permission
func (p Permission) PermissionProvider() PermissionProvider {
  return p.permissionProvider
}

func (p Permission) String() string {
  return fmt.Sprintf("Permission: Granted(%v) Role(%v)", p.granted, p.role)
}

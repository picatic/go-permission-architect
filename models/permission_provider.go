package models

import (
	perm "github.com/picatic/go-permission-architect"
)

type PermissionProvider struct {
	resourceName string
}

func NewPermissionProvider(resourceName string) perm.PermissionProvider {
	return &PermissionProvider{resourceName}
}

func (pp PermissionProvider) HandledResourceName() string {
	return pp.resourceName
}

func (pp *PermissionProvider) GetPermission(role perm.Role, permission string) perm.Permission {
	return NewPermission(permission, false, role, pp)
}

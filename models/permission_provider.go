package models

import (
	perm "github.com/picatic/go-permission-architect"
)

type PermissionProvider struct {
	resourceName string
	getPermission perm.PermissionProviderGetPermission
}

func NewPermissionProvider(resourceName string) *PermissionProvider {
	permissionProvider := new(PermissionProvider)
	permissionProvider.resourceName = resourceName
	permissionProvider.getPermission = getPermission
	return permissionProvider
}

func (pp PermissionProvider) HandledResourceName() string {
	return pp.resourceName
}

func (pp *PermissionProvider) GetPermission(role perm.Role, permission string) perm.Permission {
	return pp.getPermission(pp, role, permission)
}

func (pp *PermissionProvider) SetGetPermission(permissionProviderGetPermission perm.PermissionProviderGetPermission) {
	pp.getPermission = permissionProviderGetPermission
}

func getPermission(permissionProvider perm.PermissionProvider, role perm.Role, permission string) perm.Permission {
	return NewPermission(permission, false, role, permissionProvider)
}

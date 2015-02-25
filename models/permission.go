package models

import (
	"fmt"
	perm "github.com/picatic/go-permission-architect"
)

// Permission represents a resolved Permission
type Permission struct {
	name               string                  //name of the permission
	granted            bool                    //if the permission was granted or not
	role               perm.Role               //the role used to generate the permission
	permissionProvider perm.PermissionProvider //reference to the PermissionProvider used to create this Permission
}

//NewPermission creates a new Permission model
func NewPermission(name string, granted bool, role perm.Role, provider perm.PermissionProvider) perm.Permission {
	return &Permission{name, granted, role, provider}
}

//PermissionName returns the name
func (p Permission) PermissionName() string {
	return p.name
}

//Granted returns the boolean value indicating if this permission is granted
func (p Permission) Granted() bool {
	return p.granted
}

//SetGranted sets the granted status of the permission
func (p *Permission) SetGranted(granted bool) {
	p.granted = granted
}

//Role returns the Role used to resolve this Permission
func (p Permission) Role() perm.Role {
	return p.role
}

//SetRole sets the Role on the Permission
func (p *Permission) SetRole(role perm.Role) {
	p.role = role
}

//PermissionProvider is a reference to the PermissionProvider that generated this Permission
func (p Permission) PermissionProvider() perm.PermissionProvider {
	return p.permissionProvider
}

//SetPermissionProvider sets the PermissionProvider
func (p *Permission) SetPermissionProvider(pp perm.PermissionProvider) {
	p.permissionProvider = pp
}

func (p Permission) String() string {
	return fmt.Sprintf("Permission[%s] Granted(%v) %v", p.name, p.granted, p.role)
}

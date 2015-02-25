package permission

import (
	"log"
)

type Session interface {
	Name() string
	SetName(name string) 

	RoleProviders() []RoleProvider
	RegisterRoleProvider(roleProvider RoleProvider) error
	RoleProviderFor(profileName string, resourceName string) RoleProvider

	PermissionProviders() []PermissionProvider
	RegisterPermissionProvider(permissionProvider PermissionProvider) error
	PermissionProviderFor(resourceName string) PermissionProvider

	Logger() *log.Logger
	SetLogger(logger *log.Logger)

	DefaultRole(profile Profile, resource Resource) Role
	GetRole(p Profile, r Resource) Role
	GetPermission(p Profile, r Resource, permission string) Permission
}

//Profile interface represents a requesting user, group, organizational unit, etc.
type Profile interface {
	ProfileName() string
	ProfileIdentifier() string
}

//Resource interface represents something that can have permissions
type Resource interface {
	ResourceName() string
	ResourceIdentifier() string
}

//Role interface represents a role relationalship between a Profile and Resource
type Role interface {
	RoleName() string
	Profile() Profile
	SetProfile(Profile)
	Resource() Resource
	SetResource(Resource)
	RoleProvider() RoleProvider
	SetRoleProvider(RoleProvider)
}

//RoleProvider provides an interface to ask what role or roles a Profile and Resource matching would have
type RoleProvider interface {
	HandledProfileName() string
	HandledResourceName() string
	AllRoles(p Profile, r Resource) []Role //Returns all the applicable roles a Profile and Resource could potentially have. Ordered by
	BestRole(p Profile, r Resource) Role
}

//Permission represents the answer to "Does Role with Resource have this `permission`?"
type Permission interface {
	PermissionName() string
	Granted() bool
	SetGranted(bool)
	Role() Role
	SetRole(Role)
	PermissionProvider() PermissionProvider
	SetPermissionProvider(PermissionProvider)
}

//PermissionProvider
type PermissionProvider interface {
	HandledResourceName() string
	GetPermission(role Role, permission string) Permission
}

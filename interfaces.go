package permission

import (
	"log"
)

//Session interface hooks up registering Providers and answering questions about Roles and Permissions
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

	SetContext(context interface{})
	Context() interface{}

	NewSession(name string) Session
	SetParent(sess Session)
	Parent() Session
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

type RoleProviderAllRoles func(roleProvider RoleProvider, p Profile, r Resource) []Role
type RoleProviderBestRole func(ropeProvider RoleProvider, p Profile, r Resource) Role

//RoleProvider provides an interface to ask what role or roles a Profile and Resource matching would have
type RoleProvider interface {
	HandledProfileName() string
	HandledResourceName() string
	AllRoles(profile Profile, resource Resource) []Role // (p Profile, r Resource) []Role //Returns all the applicable roles a Profile and Resource could potentially have. Ordered by
	SetAllRoles(roleProviderAllRoles RoleProviderAllRoles)
	BestRole(p Profile, r Resource) Role
	SetBestRole(roleProviderBestRole RoleProviderBestRole)
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

type PermissionProviderGetPermission func(permissionProvider PermissionProvider, role Role, permission string) Permission

//PermissionProvider
type PermissionProvider interface {
	HandledResourceName() string
	GetPermission(role Role, permission string) Permission
	SetGetPermission(getPermission PermissionProviderGetPermission)
}

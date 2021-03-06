package permission

import (
	"log"
	"os"
)

//session provides a named space to register your PermissionProviders and RoleProviders.
//Provided as a base implementation that should work for most cases.
type session struct {
	name                string               // name of this Session
	roleProviders       []RoleProvider       // Registered RoleProviders
	permissionProviders []PermissionProvider // Registered PermissionProviders
	defaultRoleFunc     DefaultRoleFunc      // Function that can resolve a default role
	logger              *log.Logger          // Logger to use for errors/warnings/info
	ctx                 interface{}          // A user assigned context for this session
	parent              Session              // if not nil, this is the parent Session
}

type DefaultRoleFunc func(session Session, profile Profile, resource Resource) (Role, error)

//newSession creates a new Session for registering Role and Permission Providers
func newSession(name string) Session {
	s := new(session)
	s.name = name
	s.roleProviders = []RoleProvider{}
	s.permissionProviders = []PermissionProvider{}
	s.defaultRoleFunc = nil
	s.logger = log.New(os.Stdout, "PERMISSION: ", log.LstdFlags)
	return s
}

//Name gets the name of the session
func (s session) Name() string {
	return s.name
}

//SetName set the name of the session
func (s *session) SetName(name string) {
	s.name = name
}

//SetLogger sets a logger to use
func (s *session) SetLogger(logger *log.Logger) {
	s.logger = logger
}

//Logger returns the logger
func (s session) Logger() *log.Logger {
	return s.logger
}

//RoleProviders returns the registered RoleProviders with this session
func (s session) RoleProviders() []RoleProvider {
	return s.roleProviders
}

//RegisterRoleProvider registers a unique RoleProvider
func (s *session) RegisterRoleProvider(roleProvider RoleProvider) error {
	rp, _ := s.RoleProviderFor(roleProvider.HandledProfileName(), roleProvider.HandledResourceName())
	if rp != nil {
		return NewDuplicateRoleProviderError(roleProvider.HandledProfileName(), roleProvider.HandledResourceName())
	}
	roleProvider.SetSession(s)
	s.roleProviders = append(s.roleProviders, roleProvider)
	return nil
}

//RoleProviderFor returns a matching RoleProvider based on Profile and Resource model names
func (s session) RoleProviderFor(profileName string, resourceName string) (RoleProvider, error) {
	for i := range s.roleProviders {
		if s.roleProviders[i].HandledProfileName() == profileName &&
			s.roleProviders[i].HandledResourceName() == resourceName {
			return s.roleProviders[i], nil
		}
	}
	if s.parent != nil {
		return s.parent.RoleProviderFor(profileName, resourceName)
	}
	return nil, NewRoleProviderNotFoundError(profileName, resourceName)
}

//PermissionProviders returns registered PermissionProviders
func (s session) PermissionProviders() []PermissionProvider {
	return s.permissionProviders
}

//RegisterPermissionProvider registers a unique PermissionProvider with the session
func (s *session) RegisterPermissionProvider(permissionProvider PermissionProvider) error {
	pp, _ := s.PermissionProviderFor(permissionProvider.HandledResourceName())
	if pp != nil {
		return NewDuplicatePermisionProviderError(permissionProvider.HandledResourceName())
	}
	permissionProvider.SetSession(s)
	s.permissionProviders = append(s.permissionProviders, permissionProvider)
	return nil
}

//PermissionProviderFor returns a PermissionProvider for the given Resource model name
func (s session) PermissionProviderFor(resourceName string) (PermissionProvider, error) {
	for i := range s.permissionProviders {
		if s.permissionProviders[i].HandledResourceName() == resourceName {
			return s.permissionProviders[i], nil
		}
	}
	if s.parent != nil {
		return s.parent.PermissionProviderFor(resourceName)
	}
	return nil, NewPermissionProviderNotFound(resourceName)
}

//DefaultRole returns the default Role when all else fails to resolve
func (s *session) DefaultRole(profile Profile, resource Resource) (Role, error) {
	if s.defaultRoleFunc != nil {
		return s.defaultRoleFunc(s, profile, resource)
	}

	return nil, DefaultRoleNotSetError
}

func (s *session) SetDefaultRole(fn DefaultRoleFunc) {
	s.defaultRoleFunc = fn
}

//GetRole returns what role a Profile and Resource have
func (s session) GetRole(profile Profile, resource Resource) (Role, error) {
	roleProvider, err := s.RoleProviderFor(profile.ProfileName(), resource.ResourceName())
	if err != nil {
		return nil, err
	}
	if roleProvider == nil {
		role, _ := s.DefaultRole(profile, resource)
		return role, nil
	}
	return roleProvider.BestRole(profile, resource)
}

//GetPermission
func (s session) GetPermission(profile Profile, resource Resource, permission string) (Permission, error) {
	role, err := s.GetRole(profile, resource)
	if err != nil {
		return nil, err
	}
	permissionProvider, err := s.PermissionProviderFor(resource.ResourceName())
	if err != nil {
		return nil, err
	}
	return permissionProvider.GetPermission(role, permission)
}

// SetContext for this Session
func (s *session) SetContext(context interface{}) {
	s.ctx = context
}

// Context get the context assigned to this Session
func (s session) Context() interface{} {
	return s.ctx
}

// NewSession as a child of this Session
func (s *session) NewSession(name string) Session {
	child := newSession(name)
	child.SetParent(s)
	return child
}

// SetParent of this Session
func (s *session) SetParent(sess Session) {
	s.parent = sess
}

// Parent of this Session
func (s session) Parent() Session {
	return s.parent
}

var sessions = map[string]Session{} //Singleton registry of Sessions

//GetSession returns an existing named Session or creates a new Session if it does not exist
func GetSession(name string) Session {
	if val, ok := sessions[name]; ok {
		return val
	} else {
		sessions[name] = newSession(name)
		return sessions[name]
	}
}

//RegisterSession register a session within the singleton
func RegisterSession(session Session) {
	sessions[session.Name()] = session
}

package permission

import (
	"log"
	"os"
)

//Session provides a named space to register your PermissionProviders and RoleProviders
type session struct {
	name                string               // name of this Session
	roleProviders       []RoleProvider       // Registered RoleProviders
	permissionProviders []PermissionProvider // Registered PermissionProviders
	defaultRole         Role
	logger              *log.Logger
}

//newSession creates a new Session for registering Role and Permission Providers
func newSession(name string) Session {
	s := new(session)
	s.name = name
	s.roleProviders = []RoleProvider{}
	s.permissionProviders = []PermissionProvider{}
	s.logger = log.New(os.Stdout, "PERMISSION: ", log.LstdFlags)
	return s
}

func (s session) Name() string {
	return s.name
}

func (s *session) SetName(name string) {
	s.name = name
}

func (s *session) SetLogger(logger *log.Logger) {
	s.logger = logger
}

func (s session) Logger() *log.Logger {
	return s.logger
}

func (s session) RoleProviders() []RoleProvider {
	return s.roleProviders
}

func (s session) RegisterRoleProvider(roleProvider RoleProvider) {
	s.roleProviders = append(s.roleProviders, roleProvider)
}

func (s session) RoleProviderFor(profileName string, resourceName string) RoleProvider {
	return nil
}

func (s session) PermissionProviders() []PermissionProvider {
	return s.permissionProviders
}

func (s session) RegisterPermissionProvider(permissionProvider PermissionProvider) {
	s.permissionProviders = append(s.permissionProviders, permissionProvider)
}

func (s session) PermissionProviderFor(permission string) PermissionProvider {
	return nil
}

func (s session) DefaultRole() Role {
	return s.defaultRole
}

//GetRole returns what role a Profile and Resource have
func (s session) GetRole(profile Profile, resource Resource) Role {
	return nil
}

func (s session) GetPermission(profile Profile, resource Resource, permission string) Permission {
	return nil
}

var sessions = map[string]Session{}

//GetSession returns an existing named Session or creates a new Session if it does not exist
func GetSession(name string) Session {
	if val, ok := sessions[name]; ok {
		return val
	} else {
		sessions[name] = newSession(name)
		return sessions[name]
	}
}

func RegisterSession(session Session) {
	sessions[session.Name()] = session
}




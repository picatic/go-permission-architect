package permission

//Session provides a named space to register your PermissionProviders and RoleProviders
type Session struct {
  Name string // name of this Session
  RoleProviders []RoleProvider // Registered RoleProviders
  PermissionProviders []PermissionProvider // Registered PermissionProviders
}

func (s *Session)RegisterRoleProvider(rp RoleProvider) {
  s.RoleProviders = append(s.RoleProviders, rp)
}

//newSession creates a new Session for registering Role and Permission Providers
func newSession(name string) *Session {
  s := new(Session)
  s.Name = name
  s.RoleProviders = []RoleProvider{}
  s.PermissionProviders = []PermissionProvider{}
  return s
  //return &Session{name, []RoleProvider{}, []PermissionProvider{}}
}

var sessions = map[string]*Session{}

//GetSession returns an existing named Session or creates a new Session if it does not exist
func GetSession(name string) *Session {
  if val, ok := sessions[name]; ok {
    return val
  } else {
    sessions[name] = newSession(name)
    return sessions[name]
  }
}

// func (s *Session) GetRole(profile Profile, resource, Resource) Role {

// }

// func (s *Session) GetPermission(profile Profile, resource Resource, permission string) Permission {

// }

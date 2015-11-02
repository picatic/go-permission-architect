package models

import (
	"fmt"
	perm "github.com/picatic/go-permission-architect"
)

//RoleProvider base implementation, useful for building on top of, but not much else
type RoleProvider struct {
	handledProfileName  string
	handledResourceName string
	allRoles            perm.RoleProviderAllRoles
	bestRole            perm.RoleProviderBestRole
	session             perm.Session
}

//NewRoleProvider create a new RoleProvider that handles Profile and Resource by name
func NewRoleProvider(profileName string, resourceName string) perm.RoleProvider {
	rp := new(RoleProvider)
	rp.handledProfileName = profileName
	rp.handledResourceName = resourceName
	rp.allRoles = allRoles
	rp.bestRole = bestRole
	return rp
}

//AllRoles returns all applicable roles, but this implementation just returns one role of guest
func (rp *RoleProvider) AllRoles(profile perm.Profile, resource perm.Resource) []perm.Role {
	return rp.allRoles(rp, profile, resource)
}

func allRoles(rp perm.RoleProvider, profile perm.Profile, resource perm.Resource) []perm.Role {
	var roles []perm.Role
	roles = append(roles, NewRole("guest", profile, resource, rp))
	return roles
}

func (rp *RoleProvider) SetAllRoles(roleProviderAllRoles perm.RoleProviderAllRoles) {
	rp.allRoles = roleProviderAllRoles
}

//BestRole returns the best role, usual the first Role from AllRoles
func (rp *RoleProvider) BestRole(p perm.Profile, r perm.Resource) perm.Role {
	return rp.bestRole(rp, p, r)
}

func bestRole(roleProvider perm.RoleProvider, p perm.Profile, r perm.Resource) perm.Role {
	roles := roleProvider.AllRoles(p, r)
	return bestRoleWithRoles(roleProvider, p, r, roles)
}

func bestRoleWithRoles(roleProvider perm.RoleProvider, p perm.Profile, r perm.Resource, roles []perm.Role) perm.Role {
	if len(roles) >= 1 {
		return roles[0]
	} else {
		return NewRole("guest", p, r, roleProvider)
	}
}

func (rp *RoleProvider) SetBestRole(roleProviderBestRole perm.RoleProviderBestRole) {
	rp.bestRole = roleProviderBestRole
}

//HandledProfileName indicates what Profile name this RoleProvider handles
func (rp RoleProvider) HandledProfileName() string {
	return rp.handledProfileName
}

//HandledResourceName indicates what Resource name this RoleProvider handles
func (rp *RoleProvider) HandledResourceName() string {
	return rp.handledResourceName
}

//String makes a pretty string
func (rp RoleProvider) String() string {
	return fmt.Sprintf("RoleProvider[%s][%s]", rp.handledProfileName, rp.handledResourceName)
}

//SetSession sets the Session on this RoleProvider
func (rp *RoleProvider) SetSession(sess perm.Session) {
	rp.session = sess
}

//Session returns the Session this RoleProvider was initialized with
func (rp RoleProvider) Session() perm.Session {
	return rp.session
}

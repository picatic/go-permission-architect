package models

import (
	"fmt"
	perm "github.com/picatic/go-permission-architect"
)

// Profile represents a requesting identity. This could be a User, Group, Service or whatever would request a Permission
type Profile struct {
	name       string
	identifier string
}

func NewProfile(name string, identifier string) perm.Profile {
	return &Profile{name, identifier}
}

func (p Profile) ProfileName() string {
	return p.name
}

func (p Profile) ProfileIdentifier() string {
	return p.identifier
}

func (p Profile) String() string {
	return fmt.Sprintf("Profile[%s][%s]", p.name, p.identifier)
}

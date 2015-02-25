package models

import (
	"fmt"
	perm "github.com/picatic/go-permission-architect"
)

//Resource represents a Resource model which will have permission associated to it via a PermissionProvider.
type Resource struct {
	name       string
	identifier string
}

//NewResource creates a new Resource model
func NewResource(name string, identifier string) perm.Resource {
	return &Resource{name, identifier}
}

//ResourceName returns the name of the model
func (r Resource) ResourceName() string {
	return r.name
}

//ResourceIdentifier retursn the identifier for the model
func (r Resource) ResourceIdentifier() string {
	return r.identifier
}

func (r Resource) String() string {
	return fmt.Sprintf("Resource[%s][%s]", r.name, r.identifier)
}

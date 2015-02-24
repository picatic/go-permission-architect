package models

import (
  "fmt"
  perm "github.com/picatic/go-permission-architect"
)

type Resource struct {
  name string
  identifier string
}

func NewResource(name string, identifier string) perm.Resource {
  return &Resource{name, identifier}
}

func (r Resource) ResourceName() string {
  return r.name
}

func (r Resource) ResourceIdentifier() string {
  return r.identifier
}

func (r Resource) String() string {
  return fmt.Sprintf("Resource[%s][%s]", r.name, r.identifier)
}

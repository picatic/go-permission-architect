package models

import (
  "fmt"
)

type Resource struct {
  name string
  identifier string
}

func NewResource(name string, identifier string) Resource {
  return Resource{name, identifier}
}

func (r Resource) Name() string {
  return r.name
}

func (r Resource) Identifier() string {
  return r.identifier
}

func (r Resource) String() string {
  return fmt.Sprintf("Resource[%s][%s]", r.name, r.identifier)
}

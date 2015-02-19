package models

import (
  "testing"
  "fmt"
)

func TestNewPermission(t *testing.T) {
  role := NewRole()
  permissionProvider := NewPermissionProvider()
  p := NewPermission(true, role, permissionProvider)
  if (p.Granted() != true) {
    t.Fail()
  }
  if (p.Role() != role) {
    t.Fail()
  }
  if (p.PermissionProvider() != permissionProvider) {
    t.Fail()
  }
}

func ExampleNewPermission() {
  permission := NewPermission(true, Role{}, PermissionProvider{})
  fmt.Println(permission)
  //Output: Permission: Granted(true) Role({})
}

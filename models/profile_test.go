package models

import (
  "testing"
  "fmt"
)

func TestNewProfile(t *testing.T) {
  p := NewProfile("User", "1")
  if (p.Name() != "User") {
    t.Fail()
  }
  if (p.Identifier() != "1") {
    t.Fail()
  }
}

func ExampleProfile() {
  p := NewProfile("Ex", "123")
  fmt.Println(p)
  //Output: Profile[Ex][123]
}

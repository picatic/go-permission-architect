package models

import (
  "testing"
  "fmt"
)

func TestNewResource(t *testing.T) {
  p := NewResource("Post", "2")
  if (p.Name() != "Post") {
    t.Fail()
  }
  if (p.Identifier() != "2") {
    t.Fail()
  }
}

func ExampleResource() {
  p := NewResource("Post", "3")
  fmt.Println(p)
  //Output: Resource[Post][3]
}

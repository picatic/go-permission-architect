package permission

import (
	"errors"
	"fmt"
)

var (
	DefaultRoleNotSetError = errors.New("DefaultRole not set on Session")
)

type RoleProviderNotFoundError struct {
	Profile  string
	Resource string
}

func NewRoleProviderNotFoundError(profile string, resource string) error {
	return &RoleProviderNotFoundError{profile, resource}
}

func (e RoleProviderNotFoundError) Error() string {
	return fmt.Sprintf("Role Provider Not Found %s %s", e.Profile, e.Resource)
}

type DuplicateRoleProviderError struct {
	Profile  string
	Resource string
}

func NewDuplicateRoleProviderError(profile string, resource string) error {
	return &DuplicateRoleProviderError{profile, resource}
}

func (e DuplicateRoleProviderError) Error() string {
	return fmt.Sprintf("Duplicate Role Provider %s %s", e.Profile, e.Resource)
}

type PermissionProviderNotFoundError struct {
	Resource string
}

func NewPermissionProviderNotFound(resource string) error {
	return &PermissionProviderNotFoundError{resource}
}

func (e PermissionProviderNotFoundError) Error() string {
	return fmt.Sprintf("Duplicate Permission Provider %s", e.Resource)
}

type DuplicatePermissionProviderError struct {
	Resource string
}

func NewDuplicatePermisionProviderError(resource string) error {
	return &DuplicatePermissionProviderError{resource}
}

func (e DuplicatePermissionProviderError) Error() string {
	return fmt.Sprintf("Permission Provider Not Found %s", e.Resource)
}

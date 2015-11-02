package example

import (
	session "github.com/picatic/go-permission-architect"
	models "github.com/picatic/go-permission-architect/models"
)

type PostPermissionProvider struct {
	*models.PermissionProvider
}

func newPostPermissionProvider() *PostPermissionProvider {
	ppp := &PostPermissionProvider{models.NewPermissionProvider("Post")}
	return ppp
}

func (ppp *PostPermissionProvider) GetPermission(role session.Role, permission string) (session.Permission, error) {
	if role.RoleName() == "admin" {
		return models.NewPermission(permission, true, role, ppp), nil
	} else {
		return models.NewPermission(permission, true, role, ppp), nil
	}

}

func init() {
	session.GetSession("default").RegisterPermissionProvider(newPostPermissionProvider())
}

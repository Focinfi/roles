package roles

import (
	_ "fmt"
	"github.com/Focinfi/gset"
)

// Roler assume that a user implemented Role() to let this roles package
// learn roles it has
type Roler interface {
	Roles() []string
}

// Resourcer assume that a resource model implemented TableName() to let roles
// package learn which kind of this reasource is.
type Resourcer interface {
	TableName() string
}

// roles contains all roles in a map
var roles = make(map[string]*Role)

type Role struct {
	allowPermissions map[string]*gset.Set
}

// Allow add perssion set into a role
func (r *Role) Allow(resourcer Resourcer, permissions ...Permission) *Role {
	permissionsE := make([]gset.IdGetter, len(permissions))
	for i, permission := range permissions {
		permissionsE[i] = permission
	}

	permissionSet, ok := r.allowPermissions[resourcer.TableName()]
	if ok {
		permissionSet.Add(permissionsE...)
	} else {
		r.allowPermissions[resourcer.TableName()] = gset.NewSet(permissionsE...)
	}

	return r
}

// NewRole return a pointer of a new Role
func NewRole() *Role {
	return &Role{make(map[string]*gset.Set)}
}

// Add add new role into roles
func Add(name string) *Role {
	role, ok := roles[name]
	if !ok {
		role = NewRole()
		roles[name] = role
	}
	return role
}

// Can check roler's permission of the resourcer
func Can(roler Roler, resourcer Resourcer, permission Permission) bool {
	roleNames := roler.Roles()
	for _, roleName := range roleNames {
		if roles[roleName].allowPermissions[resourcer.TableName()].Has(permission) {
			return true
		}
	}
	return false
}

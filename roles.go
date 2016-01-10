package roles

import (
	_ "fmt"
	"github.com/Focinfi/gset"
)

type Roler interface {
	Roles() []string
}

type Resourcer interface {
	TableName() string
}

var roles = make(map[string]*Role)

type Role struct {
	allowPermissions map[string]*gset.Set
}

func (r *Role) Allow(resourcer Resourcer, permissions ...gset.Elementer) *Role {
	permissionSet, ok := r.allowPermissions[resourcer.TableName()]
	if ok {
		permissionSet.Add(permissions...)
	} else {
		r.allowPermissions[resourcer.TableName()] = gset.NewSet(permissions...)
	}
	return r
}

func NewRole() *Role {
	return &Role{make(map[string]*gset.Set)}
}

func Add(name string) *Role {
	role, ok := roles[name]
	if !ok {
		role = NewRole()
		roles[name] = role
	}
	return role
}

func Can(roler Roler, resourcer Resourcer, permission Permission) bool {
	roleNames := roler.Roles()
	for _, roleName := range roleNames {
		if roles[roleName].allowPermissions[resourcer.TableName()].Has(permission) {
			return true
		}
	}
	return false
}

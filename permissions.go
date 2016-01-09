package roles

import (
	"sync"
)

var mux sync.RWMutex

type Permission int

const (
	Read Permission = iota + 1
	Create
	Update
	Delete
	CRUD
)

func (p Permission) Element() interface{} {
	return p
}

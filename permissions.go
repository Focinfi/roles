package roles

type Permission int

const (
	Read Permission = iota + 1
	Create
	Update
	Delete
	CRUD
)

// Element keeps to the Elementer interface
func (p Permission) Element() interface{} {
	return p
}

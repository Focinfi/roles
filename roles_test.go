package roles

import (
	"testing"
)

type Book struct{}

func (b Book) TableName() string {
	return "books"
}

type User struct{}

func (u User) Roles() []string {
	return []string{"admin"}
}

func TestAdd(t *testing.T) {
	Add("admin")
	if _, ok := roles["admin"]; !ok {
		t.Error("Can not add admin role")
	}
}

func TestAllow(t *testing.T) {
	Add("admin")
	book := Book{}
	roles["admin"].Allow(book, CRUD)
	if !roles["admin"].allowPermissions[book.TableName()].Has(CRUD) {
		t.Error("can not allow a admin to crud user")
	}
}

func TestCan(t *testing.T) {
	Add("admin")
	book := Book{}
	user := User{}
	roles["admin"].Allow(book, Read)
	if !Can(user, book, Read) {
		t.Error("can not check if a user can read books")
	}
}

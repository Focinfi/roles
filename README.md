Resources access manangment for different roles

### Install 

`go get github.com/Focinfi/roles`


### PermissioMode
  1. roles.Read
  2. roles.Create
  3. roles.Update
  4. roles.Delete
  5. roles.CURD

### Roler interface

Implement the Roler interface for you model struct, like User.

```go
  type Roler interface {
    Roles() []string
  }
```

#### Resourcer interface

Implement the Rescource interface for your Resource struct, like Book.

```go
  type Resourcer interface {
    TableName() string
  }

```
  
### API

#### Add a role

```go
 visitorRole := roles.Add("visitor")
```

#### Allow a role has some permissions

```go
  // Here you have a Book struct
  type Book struct{}
  // Implement the Resourcer interface
  func (b Book) TableName() string {
    return "books"
  }

  // Then you can all visitors can only read books
  visitorRole.Allow(Book{}, roles.Read)
```

#### Check if a role's access of a book

```go
  // Here you have a User struct
  type User struct{}

  // Implemetn the Roler interface
  func (u User) Roles() []string {
    // Query from database and get the roles array
    return roles 
  }

  if roels.Can(User{}, Book{}, roles.Read) {
    // do something when he can read a book
  } else {
    // do something when he can not read a book
  }
```


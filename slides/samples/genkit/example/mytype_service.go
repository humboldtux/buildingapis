// generated by genkit -- DO NOT EDIT
package example

import (
	"errors"

	"github.com/m4rw3r/uuid"
)


// UserService represents operations on a User
type UserService interface {
	Create(User) (string, error)
	Get(string) (User, error)
	Update(User) error
	List() ([]User, error)
	Delete(string) error
}

func (t User ) String() string {
	return t.ID.String()
}

type userService struct {
	userList map[string]User
}

func (t userService) Create(c User) (string, error) {
	// Create a member here
	u, _ := uuid.V4()
	c.ID = u
	t.userList[u.String()] = c
	return u.String(), nil
}
func (t userService)  Get(id string) (User, error) {
	// retrieve user
	member, ok := t.userList[id]
	if !ok {
		return member, ErrNotFound
	}
	return member, nil
}
func (t userService) Update(i User) error {
	// update
	t.userList[i.ID.String()] = i
	return nil
}
func (t userService)  List() ([]User, error) {
	// get all
	return []User{}, nil
}
func (t userService) Delete(id string) error {
	// delete user
	delete(t.userList, id)
	return nil
}

// ErrExists is returned when the user already exists
var ErrExists = errors.New("User Exists")
var ErrNotFound = errors.New("User Not Found")

// ServiceMiddleware is a chainable behavior modifier for UserService.
type ServiceMiddleware func(UserService) UserService
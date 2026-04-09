package user

import (
	shared2 "github.com/oderapi/src/domain/shared"
)

type User struct {
	id       shared2.ID
	name     string
	email    string
	username string
	password string
	status   Status
	roles    []Role
}

func NewFirstSuperAdmin(id shared2.ID, name string, email string, username string, password string) (User, *shared2.DomainError) {
	v := &shared2.Validator{}

	name = v.Check(name, validateName)
	email = v.Check(email, validateEmail)
	username = v.Check(username, validateUsername)
	password = v.Check(password, validatePassword)

	if v.Err != nil {
		return User{}, v.Err
	}

	roles := []Role{SUPER_ADMIN}
	return User{id: id, name: name, email: email, username: username, password: password, status: ACTIVE, roles: roles}, nil
}

func With(id shared2.ID, name string, email string, username string, password string, status Status, roles []Role) User {
	return User{id: id, name: name, email: email, username: username, password: password, status: status, roles: roles}

}

func (u User) Id() shared2.ID   { return u.id }
func (u User) Name() string     { return u.name }
func (u User) Email() string    { return u.email }
func (u User) Username() string { return u.username }
func (u User) Password() string { return u.password }
func (u User) Status() Status   { return u.status }
func (u User) Roles() []Role    { return u.roles }

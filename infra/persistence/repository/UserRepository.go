package repository

import "github.com/oderapi/domain/entity/user"

type UserRepository interface {
	ExistsActiveSuperAdmin() (bool, error)
	Save(user user.User) error
}

package repository

import (
	"github.com/oderapi/src/domain/entity/user"
)

type UserRepository interface {
	ExistsActiveSuperAdmin() (bool, error)
	Save(user user.User) error
}

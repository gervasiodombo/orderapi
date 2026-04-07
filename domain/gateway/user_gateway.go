package gateway

import (
	"github.com/oderapi/domain/entity/user"
)

type UserGateway interface {
	ExistsActiveSuperAdmin() (bool, error)
	Save(user user.User) error
}

package gateway

import (
	"github.com/oderapi/domain/entity/user"
)

type UserGateway interface {
	ExistsActiveSuperAdmin(username string) bool
	Save(user user.User) error
}

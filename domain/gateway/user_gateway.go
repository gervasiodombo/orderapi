package gateway

import (
	"github.com/oderapi/domain/entity/user"
	"github.com/oderapi/domain/vo"
)

type UserGateway interface {
	ExistsActiveSuperAdmin(username string) bool
	Save(user user.User) (*vo.Output, error)
}

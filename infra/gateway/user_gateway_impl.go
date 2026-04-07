package gateway

import (
	"github.com/oderapi/domain/entity/user"
	"github.com/oderapi/infra/persistence/repository"
)

type UserGatewayImpl struct {
	repository repository.UserRepositoryImpl
}

func NewUserGatewayImpl() *UserGatewayImpl {
	return &UserGatewayImpl{}
}

func (ug *UserGatewayImpl) ExistsActiveSuperAdmin() (bool, error) {
	return ug.repository.ExistsActiveSuperAdmin()
}

func (ug *UserGatewayImpl) Save(user user.User) error {
	return ug.repository.Save(user)
}

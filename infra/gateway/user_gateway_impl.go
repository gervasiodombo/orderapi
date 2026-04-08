package gateway

import (
	"github.com/oderapi/domain/entity/user"
	"github.com/oderapi/infra/persistence/repository"
)

type UserGatewayImpl struct {
	userRepository repository.UserRepository
}

func NewUserGatewayImpl(userRepository repository.UserRepository) *UserGatewayImpl {
	return &UserGatewayImpl{userRepository: userRepository}
}

func (ug *UserGatewayImpl) ExistsActiveSuperAdmin() (bool, error) {
	return ug.userRepository.ExistsActiveSuperAdmin()
}

func (ug *UserGatewayImpl) Save(user user.User) error {
	return ug.userRepository.Save(user)
}

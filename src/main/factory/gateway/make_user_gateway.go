package gateway

import (
	"github.com/oderapi/src/infra/gateway"
	"github.com/oderapi/src/main/factory/repository"
	"gorm.io/gorm"
)

func MakeUserGateway(postgresDb *gorm.DB) *gateway.UserGatewayImpl {
	userRepository := repository.MakeUserRepository(postgresDb)
	return gateway.NewUserGatewayImpl(userRepository)
}

package usecase

import (
	"github.com/oderapi/src/main/factory/gateway"
	"github.com/oderapi/src/usecase/user/bootstrapp_sa"
	"gorm.io/gorm"
)

func MakeBootstrapSa(postgresDb *gorm.DB) *bootstrapp_sa.BootstrapSAImpl {
	userGateway := gateway.MakeUserGateway(postgresDb)
	idGenerator := gateway.MakeIdGenerator()
	encoder := gateway.MakeEncoderGateway()
	return bootstrapp_sa.NewBootstrapSuperAdmin(userGateway, idGenerator, encoder)
}

package bootstrapp_sa

import (
	"github.com/oderapi/src/domain/entity/user"
	gateway2 "github.com/oderapi/src/domain/gateway"
	shared2 "github.com/oderapi/src/domain/shared"
	"github.com/oderapi/src/domain/vo"
)

type BootstrapSAImpl struct {
	idGenerator shared2.IDGenerator
	gateway     gateway2.UserGateway
	encoder     gateway2.EncoderGateway
}

func NewBootstrapSuperAdmin(gateway gateway2.UserGateway, idGenerator shared2.IDGenerator, encoder gateway2.EncoderGateway) *BootstrapSAImpl {
	return &BootstrapSAImpl{gateway: gateway, idGenerator: idGenerator, encoder: encoder}
}

func (b *BootstrapSAImpl) Execute(input BootstrapSAInput) (*vo.Output, *shared2.DomainError) {
	existing, err := b.gateway.ExistsActiveSuperAdmin()
	if err != nil {
		return nil, shared2.InternalError(err)
	}
	output := &vo.Output{"Super Admin successfully already created"}
	if existing {
		return output, nil
	}
	valueId := b.idGenerator.Generate()
	id, err := shared2.NewID(valueId)
	if err != nil {
		return nil, shared2.InternalError(err)
	}
	encodedPassword, err := b.encoder.Encode(input.Password)
	if err != nil {
		return nil, shared2.InternalError(err)
	}
	toSaveUser, domainErr := user.NewFirstSuperAdmin(id, input.Name, input.Email, input.Username, encodedPassword)
	if domainErr != nil {
		return nil, domainErr
	}
	err = b.gateway.Save(toSaveUser)
	if err != nil {
		return nil, shared2.InternalError(err)
	}
	output = &vo.Output{"Super Admin successfully created"}
	return output, nil
}

package bootstrapp_sa

import (
	"github.com/oderapi/domain/entity/user"
	"github.com/oderapi/domain/gateway"
	"github.com/oderapi/domain/shared"
	"github.com/oderapi/domain/vo"
)

type BootstrapSAImpl struct {
	idGenerator shared.IDGenerator
	gateway     gateway.UserGateway
	encoder     gateway.EncoderGateway
}

func NewBootstrapSuperAdmin(gateway gateway.UserGateway, idGenerator shared.IDGenerator, encoder gateway.EncoderGateway) *BootstrapSAImpl {
	return &BootstrapSAImpl{gateway: gateway, idGenerator: idGenerator, encoder: encoder}
}

func (b *BootstrapSAImpl) Execute(input BootstrapSAInput) (*vo.Output, *shared.DomainError) {
	existing, err := b.gateway.ExistsActiveSuperAdmin()
	if err != nil {
		return nil, shared.InternalError(err)
	}
	output := &vo.Output{"Super Admin successfully already created"}
	if existing {
		return output, nil
	}
	valueId := b.idGenerator.Generate()
	id, err := shared.NewID(valueId)
	if err != nil {
		return nil, shared.InternalError(err)
	}
	encodedPassword, err := b.encoder.Encode(input.Password)
	if err != nil {
		return nil, shared.InternalError(err)
	}
	toSaveUser, domainErr := user.NewFirstSuperAdmin(id, input.Name, input.Email, input.Username, encodedPassword)
	if domainErr != nil {
		return nil, domainErr
	}
	err = b.gateway.Save(toSaveUser)
	if err != nil {
		return nil, shared.InternalError(err)
	}
	output = &vo.Output{"Super Admin successfully created"}
	return output, nil
}

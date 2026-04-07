package bootstrapp_sa

import (
	"github.com/oderapi/domain/entity/user"
	"github.com/oderapi/domain/gateway"
	"github.com/oderapi/domain/shared"
	"github.com/oderapi/domain/vo"
)

type bootstrapSAImpl struct {
	idGenerator shared.IDGenerator
	gateway     gateway.UserGateway
	encoder     gateway.EncoderGateway
}

func New(gateway gateway.UserGateway, idGenerator shared.IDGenerator, encoder gateway.EncoderGateway) BootstrapSA {
	return &bootstrapSAImpl{gateway: gateway, idGenerator: idGenerator, encoder: encoder}
}

func (b *bootstrapSAImpl) Execute(input BootstrapSAInput) (*vo.Output, *shared.DomainError) {
	existing := b.gateway.ExistsActiveSuperAdmin(input.Username)
	if existing {
		return nil, nil
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
	_, err = b.gateway.Save(toSaveUser)
	if err != nil {
		return nil, shared.InternalError(err)
	}
	return nil, nil
}

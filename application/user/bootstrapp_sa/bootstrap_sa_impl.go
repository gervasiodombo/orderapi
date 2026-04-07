package bootstrapp_sa

import (
	"github.com/oderapi/domain/gateway"
	"github.com/oderapi/domain/shared"
)

type bootstrapSAImpl struct {
	idGenerator shared.IDGenerator
	gateway     gateway.UserGateway
}

func New(gateway gateway.UserGateway, idGenerator shared.IDGenerator) BootstrapSA {
	return &bootstrapSAImpl{gateway: gateway, idGenerator: idGenerator}
}

func (b *bootstrapSAImpl) Execute(input Input) (*Output, *shared.DomainError) {
	existing := b.gateway.ExistsActiveSuperAdmin(input.Username)
	if existing {
		return nil, nil
	}
	valueId := b.idGenerator.Generate()
	_, err := shared.NewID(valueId)
	if err != nil {
		return nil, shared.InternalError(err)
	}
	return nil, nil
}

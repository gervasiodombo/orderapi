package gateway

import "github.com/oderapi/domain/shared"

type EncoderGateway interface {
	encode(rawPassword string) (string, shared.DomainError)
}

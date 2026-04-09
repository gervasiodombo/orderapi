package gateway

import (
	"github.com/oderapi/src/infra/gateway"
)

func MakeEncoderGateway() *gateway.EncoderGatewayImpl {
	return gateway.NewEncoderGateway()
}

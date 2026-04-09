package gateway

import (
	"github.com/oderapi/src/infra/shared"
)

func MakeIdGenerator() *shared.UUIDGeneratorImpl {
	return shared.NewIDGenerator()
}

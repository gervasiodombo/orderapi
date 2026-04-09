package bootstrapp_sa

import (
	"github.com/oderapi/src/domain/shared"
	"github.com/oderapi/src/domain/vo"
)

type BootstrapSA interface {
	Execute(input BootstrapSAInput) (*vo.Output, *shared.DomainError)
}

package bootstrapp_sa

import (
	"github.com/oderapi/domain/shared"
	"github.com/oderapi/domain/vo"
)

type BootstrapSA interface {
	Execute(input BootstrapSAInput) (*vo.Output, *shared.DomainError)
}

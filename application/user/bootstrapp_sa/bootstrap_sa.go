package bootstrapp_sa

import "github.com/oderapi/domain/shared"

type BootstrapSA interface {
	Execute(input Input) (*Output, *shared.DomainError)
}

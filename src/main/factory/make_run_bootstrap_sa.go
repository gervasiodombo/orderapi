package factory

import (
	"github.com/oderapi/src/usecase/user/bootstrapp_sa"
)

func MakeRunBootstrapSa(input bootstrapp_sa.BootstrapSAInput, bootstrapSa *bootstrapp_sa.BootstrapSAImpl) (string, error) {
	output, err := bootstrapSa.Execute(input)
	if err != nil {
		return "", err.Cause
	}
	return output.Message, nil
}

package usecase

import (
	"os"

	"github.com/oderapi/src/usecase/user/bootstrapp_sa"
)

func MakeBootstrapSaInput() bootstrapp_sa.BootstrapSAInput {
	return bootstrapp_sa.BootstrapSAInput{
		Name:     os.Getenv("SA_NAME"),
		Email:    os.Getenv("SA_EMAIL"),
		Username: os.Getenv("SA_USERNAME"),
		Password: os.Getenv("SA_PASSWORD"),
	}
}

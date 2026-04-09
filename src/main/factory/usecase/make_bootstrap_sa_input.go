package usecase

import (
	"errors"
	"os"

	"github.com/oderapi/src/usecase/user/bootstrapp_sa"
)

func MakeBootstrapSaInput() (bootstrapp_sa.BootstrapSAInput, error) {
	name := os.Getenv("SA_NAME")
	if name == "" {
		return bootstrapp_sa.BootstrapSAInput{}, errors.New("SA_NAME environment variable not set")
	}
	email := os.Getenv("SA_EMAIL")
	if email == "" {
		return bootstrapp_sa.BootstrapSAInput{}, errors.New("SA_EMAIL environment variable not set")
	}

	username := os.Getenv("SA_USERNAME")
	if username == "" {
		return bootstrapp_sa.BootstrapSAInput{}, errors.New("SA_USERNAME environment variable not set")
	}

	password := os.Getenv("SA_PASSWORD")
	if password == "" {
		return bootstrapp_sa.BootstrapSAInput{}, errors.New("SA_PASSWORD environment variable not set")
	}

	input := bootstrapp_sa.BootstrapSAInput{
		Name:     name,
		Email:    email,
		Username: username,
		Password: password,
	}
	return input, nil
}

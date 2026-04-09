package e2e_test

import (
	"testing"

	"github.com/oderapi/src/main/factory/usecase"
	"github.com/oderapi/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestE2EShouldFailWhenEnvsMissing(t *testing.T) {
	db, cleanup := mocks.StartPostgres(t)
	defer cleanup()
	input := usecase.MakeBootstrapSaInput()
	input.Password = ""
	_, err := mocks.RunOnStart(t, db, input)
	assert.Error(t, err)
}

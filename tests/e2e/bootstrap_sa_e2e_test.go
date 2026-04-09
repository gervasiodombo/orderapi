package e2e_test

import (
	"testing"

	"github.com/oderapi/src/main/factory/usecase"
	"github.com/oderapi/src/usecase/user/bootstrapp_sa"
	"github.com/oderapi/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestE2EShouldFailWhenEnvsMissing(t *testing.T) {
	db, err := mocks.StartPostgres(t)
	defer db.Cleanup()
	input, err := usecase.MakeBootstrapSaInput()
	input.Password = ""
	_, err = mocks.RunOnStart(t, db.DB, input)
	assert.Error(t, err)
}

func TestShouldCreateSAOnFirstStartUp(t *testing.T) {
	db, err := mocks.StartPostgres(t)
	defer db.Cleanup()
	input := bootstrapp_sa.BootstrapSAInput{
		Name:     "System Admin",
		Email:    "sa@system.com",
		Username: "superadmin",
		Password: "strOnP@ssword",
	}
	output, err := mocks.RunOnStart(t, db.DB, input)
	assert.NoError(t, err)
	assert.Contains(t, output, "Super Admin successfully created")
}

func TestShouldNotCreateSAOnFirstStartUpIfAlreadyExists(t *testing.T) {
	db, err := mocks.StartPostgres(t)
	defer db.Cleanup()
	input := bootstrapp_sa.BootstrapSAInput{
		Name:     "System Admin",
		Email:    "sa@system.com",
		Username: "superadmin",
		Password: "strOnP@ssword",
	}

	output, err := mocks.RunOnStart(t, db.DB, input)
	require.NoError(t, err)

	output, err = mocks.RunOnStart(t, db.DB, input)

	assert.NoError(t, err)
	assert.Contains(t, output, "Super Admin successfully already created")
}

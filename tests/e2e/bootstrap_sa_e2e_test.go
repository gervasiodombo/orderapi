package e2e_test

import (
	"testing"

	"github.com/oderapi/src/usecase/user/bootstrapp_sa"
	"github.com/oderapi/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestE2EShouldFailWhenEnvsMissing(t *testing.T) {
	db, cleanup := mocks.StartPostgres(t)
	defer cleanup()
	input := bootstrapp_sa.BootstrapSAInput{
		Name:     "System Admin",
		Email:    "sa@system.com",
		Username: "superadmin",
	}
	_, err := mocks.RunOnStart(t, db, input)
	assert.Error(t, err)
}

func TestShouldCreateSAOnFirstStartUp(t *testing.T) {
	db, cleanup := mocks.StartPostgres(t)
	defer cleanup()
	input := bootstrapp_sa.BootstrapSAInput{
		Name:     "System Admin",
		Email:    "sa@system.com",
		Username: "superadmin",
		Password: "strOnP@ssword",
	}
	output, err := mocks.RunOnStart(t, db, input)
	assert.NoError(t, err)
	assert.Contains(t, output, "Super Admin successfully created")
}

func TestShouldNotCreateSAOnFirstStartUpIfAlreadyExists(t *testing.T) {
	db, cleanup := mocks.StartPostgres(t)
	defer cleanup()
	input := bootstrapp_sa.BootstrapSAInput{
		Name:     "System Admin",
		Email:    "sa@system.com",
		Username: "superadmin",
		Password: "strOnP@ssword",
	}

	output, err := mocks.RunOnStart(t, db, input)
	require.NoError(t, err)

	output, err = mocks.RunOnStart(t, db, input)

	assert.NoError(t, err)
	assert.Contains(t, output, "Super Admin successfully already created")
}

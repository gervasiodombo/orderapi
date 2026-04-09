package mocks

import (
	"testing"

	"github.com/oderapi/src/infra/persistence"
	"github.com/oderapi/src/main/factory"
	"github.com/oderapi/src/main/factory/usecase"
	"github.com/oderapi/src/usecase/user/bootstrapp_sa"
	"gorm.io/gorm"
)

func RunOnStart(t *testing.T, db *gorm.DB, input bootstrapp_sa.BootstrapSAInput) (string, error) {
	t.Helper()
	if err := persistence.RunMigrations(db); err != nil {
		return "", err
	}
	bootstrapSa := usecase.MakeBootstrapSa(db)
	return factory.MakeRunBootstrapSa(input, bootstrapSa)
}

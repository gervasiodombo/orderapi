package repository_test

import (
	"testing"

	"github.com/glebarez/sqlite" // ← pure Go, sem CGO
	model2 "github.com/oderapi/src/infra/persistence/model"
	"gorm.io/gorm"
)

func newTestDb(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}
	err = db.AutoMigrate(&model2.UserModel{}, &model2.RoleModel{})
	if err != nil {
		t.Fatalf("failed to migrate in-memory db: %v", err)
	}
	return db
}

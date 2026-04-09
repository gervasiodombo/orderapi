package persistence

import (
	"github.com/oderapi/infra/persistence/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.UserModel{},
		&model.RoleModel{},
	)
}

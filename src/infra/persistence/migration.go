package persistence

import (
	model2 "github.com/oderapi/src/infra/persistence/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&model2.UserModel{},
		&model2.RoleModel{},
	)
}

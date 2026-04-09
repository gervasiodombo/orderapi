package repository

import (
	"github.com/oderapi/src/infra/persistence/repository"
	"gorm.io/gorm"
)

func MakeUserRepository(postgresDb *gorm.DB) *repository.UserRepositoryImpl {
	return repository.NewUserRepositoryImpl(postgresDb)
}

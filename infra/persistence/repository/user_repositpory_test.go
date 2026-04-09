package repository_test

import (
	"testing"

	"github.com/oderapi/infra/persistence/model"
	"github.com/oderapi/infra/persistence/repository"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnFalseIfUserDoesNotExistOnExistSuperAdmin(t *testing.T) {
	//Arrange
	db := newTestDb(t)
	userRepository := repository.NewUserRepositoryImpl(db)

	//Act
	exists, err := userRepository.ExistsActiveSuperAdmin()

	//Assert
	assert.Nil(t, err)
	assert.False(t, exists)
}

func TestShouldReturnTrueIfUserDoesNotExistOnExistSuperAdmin(t *testing.T) {
	//Arrange
	db := newTestDb(t)
	userRepository := repository.NewUserRepositoryImpl(db)
	user := model.UserModel{ID: "uuid-1", Name: "Admin", Email: "sa@test.com", Username: "sa", Password: "hash", Status: "ACTIVE"}
	db.Create(&user)
	db.Create(&model.RoleModel{UserID: "uuid-1", Role: "SUPER_ADMIN"})

	//Act
	exists, err := userRepository.ExistsActiveSuperAdmin()

	//Assert
	assert.Nil(t, err)
	assert.True(t, exists)
}

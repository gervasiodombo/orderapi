package repository_test

import (
	"testing"

	"github.com/oderapi/domain/entity/user"
	"github.com/oderapi/domain/shared"
	"github.com/oderapi/infra/persistence/model"
	"github.com/oderapi/infra/persistence/repository"
	"github.com/stretchr/testify/assert"
)

// ExistsActiveSuperAdmin
func TestShouldReturnErrorIfUserDoesNotExistOnExistSuperAdminFails(t *testing.T) {
	//Arrange
	db := newTestDb(t)
	sqlDB, _ := db.DB()
	sqlDB.Close()

	userRepository := repository.NewUserRepositoryImpl(db)

	//Act
	exists, err := userRepository.ExistsActiveSuperAdmin()

	//Assert
	assert.NotNil(t, err)
	assert.False(t, exists)
	assert.Contains(t, err.Error(), "failed to check active super admin")
}

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

func TestShouldReturnErrorIfSaveFails(t *testing.T) {
	//Arrange
	db := newTestDb(t)
	sqlDB, _ := db.DB()
	sqlDB.Close()

	userRepository := repository.NewUserRepositoryImpl(db)
	id, _ := shared.NewID("test-sa-id")
	roles := []user.Role{user.SUPER_ADMIN}
	us := user.With(id, "any_name", "any_email", "any_username", "any_passq", user.ACTIVE, roles)

	//Act
	err := userRepository.Save(us)

	//Assert
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to save user")
}

func TestShouldReturnNilIfSaveSucceeds(t *testing.T) {
	//Arrange
	db := newTestDb(t)
	userRepository := repository.NewUserRepositoryImpl(db)
	id, _ := shared.NewID("test-sa-id")
	roles := []user.Role{user.SUPER_ADMIN}
	us := user.With(id, "any_name", "any_email", "any_username", "any_passq", user.ACTIVE, roles)

	//Act
	err := userRepository.Save(us)

	//Assert
	assert.Nil(t, err)
}

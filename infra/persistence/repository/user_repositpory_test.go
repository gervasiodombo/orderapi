package repository_test

import (
	"testing"

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

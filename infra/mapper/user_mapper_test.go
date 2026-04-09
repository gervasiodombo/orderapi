package mapper_test

import (
	"testing"

	"github.com/oderapi/domain/entity/user"
	"github.com/oderapi/domain/shared"
	"github.com/oderapi/infra/mapper"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnUserModelWithCorrectValues(t *testing.T) {
	//Arrange
	id, _ := shared.NewID("test-id")
	roles := []user.Role{user.SUPER_ADMIN}
	us := user.With(id, "any_name", "any_email", "any_username", "any_passq", user.ACTIVE, roles)

	//Act
	userModel := mapper.ToModel(us)

	//Assert
	assert.Equal(t, id.Value(), userModel.ID)
	assert.Equal(t, userModel.Name, userModel.Name)
	assert.Equal(t, userModel.Email, userModel.Email)
	assert.Equal(t, userModel.Username, userModel.Username)
	assert.Equal(t, userModel.Password, userModel.Password)
	assert.Equal(t, len(userModel.Roles), len(userModel.Roles))
	assert.Equal(t, userModel.Roles[0], userModel.Roles[0])
}

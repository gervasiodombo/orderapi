package mapper_test

import (
	"fmt"
	"testing"

	"github.com/oderapi/domain/entity/user"
	"github.com/oderapi/domain/shared"
	"github.com/oderapi/infra/mapper"
	"github.com/oderapi/infra/persistence/model"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnUserModelWithCorrectValues(t *testing.T) {
	//Arrange
	id, _ := shared.NewID("test-id")
	roles := []user.Role{user.SUPER_ADMIN}
	userDomain := user.With(id, "any_name", "any_email", "any_username", "any_passq", user.ACTIVE, roles)

	//Act
	userModel := mapper.ToModel(userDomain)

	//Assert
	assert.Equal(t, id.Value(), userModel.ID)
	assert.Equal(t, userDomain.Name(), userModel.Name)
	assert.Equal(t, userDomain.Email(), userModel.Email)
	assert.Equal(t, userDomain.Username(), userModel.Username)
	assert.Equal(t, userDomain.Password(), userModel.Password)
	assert.Equal(t, len(userDomain.Roles()), len(userModel.Roles))
	assert.Equal(t, userDomain.Roles()[0].String(), userModel.Roles[0].Role)
}

func TestShouldReturnErrorIfIdIsInvalidOnToDomain(t *testing.T) {
	//Arrange
	role := model.RoleModel{Role: user.SUPER_ADMIN.String()}
	roles := []model.RoleModel{role}
	userModel := model.UserModel{
		Name:     "any_name",
		Email:    "any_email",
		Username: "any_username",
		Password: "any_passq",
		Status:   user.ACTIVE.String(),
		Roles:    roles,
	}

	//Act
	_, err := mapper.ToDomain(userModel)

	//Assert
	assert.NotNil(t, err)
	assert.Contains(t, fmt.Sprint(err), "could not generate ID because value is empty")
}

func TestShouldReturnUserDomainWithCorrectValues(t *testing.T) {
	//Arrange
	id, _ := shared.NewID("test-id")
	role := model.RoleModel{UserID: id.Value(), Role: user.SUPER_ADMIN.String()}
	roles := []model.RoleModel{role}
	userModel := model.UserModel{
		ID:       "test-id",
		Name:     "any_name",
		Email:    "any_email",
		Username: "any_username",
		Password: "any_passq",
		Status:   user.ACTIVE.String(),
		Roles:    roles,
	}

	//Act
	userDomain, _ := mapper.ToDomain(userModel)

	//Assert
	assert.Equal(t, id.Value(), userModel.ID)
	assert.Equal(t, userDomain.Name(), userModel.Name)
	assert.Equal(t, userDomain.Email(), userModel.Email)
	assert.Equal(t, userDomain.Username(), userModel.Username)
	assert.Equal(t, userDomain.Password(), userModel.Password)
	assert.Equal(t, len(userDomain.Roles()), len(userModel.Roles))
	assert.Equal(t, userDomain.Roles()[0].String(), userModel.Roles[0].Role)
}

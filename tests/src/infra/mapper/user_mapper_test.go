package mapper_test

import (
	"fmt"
	"testing"

	user2 "github.com/oderapi/src/domain/entity/user"
	"github.com/oderapi/src/domain/shared"
	"github.com/oderapi/src/infra/mapper"
	model2 "github.com/oderapi/src/infra/persistence/model"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnUserModelWithCorrectValues(t *testing.T) {
	//Arrange
	id, _ := shared.NewID("test-id")
	roles := []user2.Role{user2.SUPER_ADMIN}
	userDomain := user2.With(id, "any_name", "any_email", "any_username", "any_passq", user2.ACTIVE, roles)

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
	role := model2.RoleModel{Role: user2.SUPER_ADMIN.String()}
	roles := []model2.RoleModel{role}
	userModel := model2.UserModel{
		Name:     "any_name",
		Email:    "any_email",
		Username: "any_username",
		Password: "any_passq",
		Status:   user2.ACTIVE.String(),
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
	role := model2.RoleModel{UserID: id.Value(), Role: user2.SUPER_ADMIN.String()}
	roles := []model2.RoleModel{role}
	userModel := model2.UserModel{
		ID:       "test-id",
		Name:     "any_name",
		Email:    "any_email",
		Username: "any_username",
		Password: "any_passq",
		Status:   user2.ACTIVE.String(),
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

package gateway_test

import (
	"errors"
	"testing"

	user2 "github.com/oderapi/src/domain/entity/user"
	"github.com/oderapi/src/domain/shared"
	"github.com/oderapi/src/infra/gateway"
	mocks "github.com/oderapi/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnErrorIfExistsActiveSuperAdminFails(t *testing.T) {
	//Arrange
	userRepository := &mocks.UserRepositoryMock{ExistsErr: errors.New("db connection failed")}
	userGateway := gateway.NewUserGatewayImpl(userRepository)

	//Act
	result, err := userGateway.ExistsActiveSuperAdmin()

	//Assert
	assert.NotNil(t, err)
	assert.False(t, result)
}

func TestShouldReturnBoolExistsActiveSuperAdminSucceeds(t *testing.T) {
	//Arrange
	userRepository := &mocks.UserRepositoryMock{ExistsResult: true}
	userGateway := gateway.NewUserGatewayImpl(userRepository)

	//Act
	result, err := userGateway.ExistsActiveSuperAdmin()

	//Assert
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestShouldReturnErrorIfSaveFails(t *testing.T) {
	//Arrange
	userRepository := &mocks.UserRepositoryMock{SaveErr: errors.New("db connection failed")}
	userGateway := gateway.NewUserGatewayImpl(userRepository)
	id, _ := shared.NewID("test-sa-id")
	us := user2.With(id, "any_name", "any_email", "any_username", "any_passq", user2.ACTIVE, []user2.Role{user2.SUPER_ADMIN})

	//Act
	err := userGateway.Save(us)

	//Assert
	assert.NotNil(t, err)
}

func TestShouldReturnErrorIfSaveSucceeds(t *testing.T) {
	//Arrange
	userRepository := &mocks.UserRepositoryMock{}
	userGateway := gateway.NewUserGatewayImpl(userRepository)
	id, _ := shared.NewID("test-sa-id")
	us := user2.With(id, "any_name", "any_email", "any_username", "any_passq", user2.ACTIVE, []user2.Role{user2.SUPER_ADMIN})

	//Act
	err := userGateway.Save(us)

	//Assert
	assert.Nil(t, err)
}

package gateway_test

import (
	"errors"
	"testing"

	mocks "github.com/oderapi/_mocks"
	"github.com/oderapi/infra/gateway"
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

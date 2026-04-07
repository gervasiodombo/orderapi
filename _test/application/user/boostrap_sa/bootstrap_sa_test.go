package boostrap_sa_test

import (
	"errors"
	"testing"

	"github.com/oderapi/_test/application/user/boostrap_sa/mocks"
	"github.com/oderapi/application/user/bootstrapp_sa"
	"github.com/oderapi/domain/shared"
)

func TestShouldReturnNilIfSuperAdminAlreadyExists(t *testing.T) {
	//Arrange
	gateway := &mocks.UserGatewayMock{
		SaveCalled: false,

		ExistsActiveSuperAdminResult: true,
		ExistsActiveSuperAdminCalled: true,
	}
	idGenerator := &mocks.
		IDGeneratorMock{Value: "test_sa_id"}
	usecase := bootstrapp_sa.New(gateway, idGenerator)
	input := bootstrapp_sa.Input{
		Name:     "System Admin",
		Email:    "sa@system.com",
		Password: "str0ngP@ssword",
	}

	//Act
	output, err := usecase.Execute(input)

	//Assert
	if err != nil {
		t.Errorf("should not return an error")
	}

	if output != nil {
		t.Errorf("output should be nil")
	}

	if !gateway.ExistsActiveSuperAdminResult {
		t.Error("ExistsActiveSuperAdmin should return true")
	}

	if !gateway.ExistsActiveSuperAdminCalled {
		t.Error("ExistsActiveSuperAdmin should have been called")
	}

	if gateway.SaveCalled {
		t.Error("Save() should not have been called")
	}
}

func TestShouldReturnErrorIfIdGeneratorFails(t *testing.T) {
	//Arrange
	expectedError := shared.InternalError(shared.ErrEmptyID)
	gateway := &mocks.UserGatewayMock{
		SaveCalled: false,

		ExistsActiveSuperAdminResult: false,
		ExistsActiveSuperAdminCalled: false,
	}
	idGenerator := &mocks.IDGeneratorMock{Value: ""}
	usecase := bootstrapp_sa.New(gateway, idGenerator)
	input := bootstrapp_sa.Input{
		Name:     "System Admin",
		Email:    "sa@system.com",
		Password: "str0ngP@ssword",
	}

	//Act
	output, err := usecase.Execute(input)

	//Assert
	if err == nil {
		t.Errorf("should  return an error")
	}

	if output != nil {
		t.Errorf("output should be nil")
	}

	if gateway.ExistsActiveSuperAdminResult {
		t.Error("ExistsActiveSuperAdmin should return false")
	}

	if !gateway.ExistsActiveSuperAdminCalled {
		t.Error("ExistsActiveSuperAdmin should have been called")
	}

	if gateway.SaveCalled {
		t.Error("Save() should not have been called")
	}

	if err.Code != expectedError.Code {
		t.Errorf("Should return error code: %v", expectedError.Code)
	}

	if err.Message != expectedError.Message {
		t.Errorf("Should return error message: %v", expectedError.Message)
	}

	if !errors.Is(err.Cause, expectedError.Cause) {
		t.Errorf("Should return error cause: '%v'! But received : '%v'", expectedError.Cause, err.Cause)
	}
}

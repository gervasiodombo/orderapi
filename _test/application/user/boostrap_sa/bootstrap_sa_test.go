package boostrap_sa_test

import (
	"errors"
	"testing"

	"github.com/oderapi/_test/application/user/boostrap_sa/mocks"
	"github.com/oderapi/application/user/bootstrapp_sa"
	"github.com/oderapi/domain/gateway"
	"github.com/oderapi/domain/shared"
)

func TestShouldReturnNilIfSuperAdminAlreadyExists(t *testing.T) {
	//Arrange
	username := "superAdmin"
	userGateway := &mocks.UserGatewayMock{
		ExistsActiveSuperAdminResult: true,
		ExistsActiveSuperAdminParam:  username,
	}
	idGenerator := &mocks.IDGeneratorMock{}
	encoder := &mocks.EncoderGatewayMock{}

	usecase := bootstrapp_sa.New(userGateway, idGenerator, encoder)
	input := bootstrapp_sa.Input{
		Username: username,
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

	if !userGateway.ExistsActiveSuperAdminResult {
		t.Error("ExistsActiveSuperAdmin should return true")
	}

	if !userGateway.ExistsActiveSuperAdminCalled {
		t.Error("ExistsActiveSuperAdmin should have been called")
	}

}

func TestShouldReturnErrorIfIdGeneratorFails(t *testing.T) {
	//Arrange
	expectedError := shared.InternalError(shared.ErrEmptyID)
	username := "superAdmin"
	userGateway := &mocks.UserGatewayMock{ExistsActiveSuperAdminParam: username}
	idGenerator := &mocks.IDGeneratorMock{}
	encoder := &mocks.EncoderGatewayMock{}
	usecase := bootstrapp_sa.New(userGateway, idGenerator, encoder)
	input := bootstrapp_sa.Input{
		Username: username,
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

	if userGateway.ExistsActiveSuperAdminResult {
		t.Error("ExistsActiveSuperAdmin should return false")
	}

	if !userGateway.ExistsActiveSuperAdminCalled {
		t.Error("ExistsActiveSuperAdmin should have been called")
	}

	if !idGenerator.GenerateCalled {
		t.Error("Generate should have been called")
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

func TestShouldReturnErrorIfEncoderFails(t *testing.T) {
	//Arrange
	username := "superAdmin"
	userGateway := &mocks.UserGatewayMock{ExistsActiveSuperAdminParam: username}
	idGenerator := &mocks.IDGeneratorMock{GenerateResult: "test-id"}
	encodeErr := gateway.ErrEncoder
	expectedError := shared.InternalError(encodeErr)
	encodeParam := "str0ngP@ssword"
	encoder := &mocks.EncoderGatewayMock{EncodeErr: encodeErr}

	usecase := bootstrapp_sa.New(userGateway, idGenerator, encoder)
	input := bootstrapp_sa.Input{
		Username: username,
		Name:     "System Admin",
		Email:    "sa@system.com",
		Password: encodeParam,
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

	if userGateway.ExistsActiveSuperAdminResult {
		t.Error("ExistsActiveSuperAdmin should return false")
	}

	if !userGateway.ExistsActiveSuperAdminCalled {
		t.Error("ExistsActiveSuperAdmin should have been called")
	}

	if !idGenerator.GenerateCalled {
		t.Error("Generate should have been called")
	}

	if !encoder.EncodeCalled {
		t.Error("Encoder should have been called")
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

func TestShouldReturnErrorIfNameEmpty(t *testing.T) {
	//Arrange
	username := "superAdmin"
	userGateway := &mocks.UserGatewayMock{ExistsActiveSuperAdminParam: username}
	idGenerator := &mocks.IDGeneratorMock{GenerateResult: "test-id"}
	encodeParam := "str0ngP@ssword"
	encodedPassword := "encodedPassword"
	encoder := &mocks.EncoderGatewayMock{EncodeResult: encodedPassword}
	expectedError := shared.RequiredField("User", "name")

	usecase := bootstrapp_sa.New(userGateway, idGenerator, encoder)
	input := bootstrapp_sa.Input{
		Username: username,
		Name:     "",
		Email:    "sa@system.com",
		Password: encodeParam,
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

	if userGateway.ExistsActiveSuperAdminResult {
		t.Error("ExistsActiveSuperAdmin should return false")
	}

	if !userGateway.ExistsActiveSuperAdminCalled {
		t.Error("ExistsActiveSuperAdmin should have been called")
	}

	if !idGenerator.GenerateCalled {
		t.Error("Generate should have been called")
	}

	if !encoder.EncodeCalled {
		t.Error("Encoder should have been called")
	}

	if err.Code != expectedError.Code {
		t.Errorf("Should return error code: %v", expectedError.Code)
	}

	if err.Message != expectedError.Message {
		t.Errorf("Should return error message: %v", expectedError.Message)
	}
}

func TestShouldReturnErrorIfEmailEmpty(t *testing.T) {
	//Arrange
	username := "superAdmin"
	userGateway := &mocks.UserGatewayMock{ExistsActiveSuperAdminParam: username}
	idGenerator := &mocks.IDGeneratorMock{GenerateResult: "test-id"}
	encodeParam := "str0ngP@ssword"
	encodedPassword := "encodedPassword"
	encoder := &mocks.EncoderGatewayMock{EncodeResult: encodedPassword}
	expectedError := shared.RequiredField("User", "email")

	usecase := bootstrapp_sa.New(userGateway, idGenerator, encoder)
	input := bootstrapp_sa.Input{
		Username: username,
		Name:     "System Admin",
		Email:    "",
		Password: encodeParam,
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

	if userGateway.ExistsActiveSuperAdminResult {
		t.Error("ExistsActiveSuperAdmin should return false")
	}

	if !userGateway.ExistsActiveSuperAdminCalled {
		t.Error("ExistsActiveSuperAdmin should have been called")
	}

	if !idGenerator.GenerateCalled {
		t.Error("Generate should have been called")
	}

	if !encoder.EncodeCalled {
		t.Error("Encoder should have been called")
	}

	if err.Code != expectedError.Code {
		t.Errorf("Should return error code: %v", expectedError.Code)
	}

	if err.Message != expectedError.Message {
		t.Errorf("Should return error message: %v", expectedError.Message)
	}
}

func TestShouldReturnErrorIfUsernameEmpty(t *testing.T) {
	//Arrange
	username := ""
	userGateway := &mocks.UserGatewayMock{ExistsActiveSuperAdminParam: username}
	idGenerator := &mocks.IDGeneratorMock{GenerateResult: "test-id"}
	encodeParam := "str0ngP@ssword"
	encodedPassword := "encodedPassword"
	encoder := &mocks.EncoderGatewayMock{EncodeResult: encodedPassword}
	expectedError := shared.RequiredField("User", "username")

	usecase := bootstrapp_sa.New(userGateway, idGenerator, encoder)
	input := bootstrapp_sa.Input{
		Username: username,
		Name:     "System Admin",
		Email:    "sa@system.com",
		Password: encodeParam,
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

	if userGateway.ExistsActiveSuperAdminResult {
		t.Error("ExistsActiveSuperAdmin should return false")
	}

	if !userGateway.ExistsActiveSuperAdminCalled {
		t.Error("ExistsActiveSuperAdmin should have been called")
	}

	if !idGenerator.GenerateCalled {
		t.Error("Generate should have been called")
	}

	if !encoder.EncodeCalled {
		t.Error("Encoder should have been called")
	}

	if err.Code != expectedError.Code {
		t.Errorf("Should return error code: %v", expectedError.Code)
	}

	if err.Message != expectedError.Message {
		t.Errorf("Should return error message: %v", expectedError.Message)
	}
}

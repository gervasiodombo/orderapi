package bootstrap_sa_test

import (
	"errors"
	"testing"

	shared2 "github.com/oderapi/src/domain/shared"
	bootstrapp_sa2 "github.com/oderapi/src/usecase/user/bootstrapp_sa"
	"github.com/oderapi/tests/mocks"
)

func TestShouldReturnErrorIfExistsVerificationFails(t *testing.T) {
	//Arrange
	username := "superAdmin"
	expectedError := errors.New("verification failed")
	userGateway := &mocks.UserGatewayMock{ExistsActiveSuperAdminErr: expectedError}
	idGenerator := &mocks.IDGeneratorMock{}
	encoder := &mocks.EncoderGatewayMock{}

	usecase := bootstrapp_sa2.NewBootstrapSuperAdmin(userGateway, idGenerator, encoder)
	input := bootstrapp_sa2.BootstrapSAInput{
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

	if err.Cause != expectedError {
		t.Errorf("Should return error message: %v", expectedError)
	}
}

func TestShouldReturnOutPutMessageIFfSuperAdminAlreadyExists(t *testing.T) {
	//Arrange
	username := "superAdmin"
	userGateway := &mocks.UserGatewayMock{ExistsActiveSuperAdminResult: true}
	idGenerator := &mocks.IDGeneratorMock{}
	encoder := &mocks.EncoderGatewayMock{}
	expectedOutput := "Super Admin successfully already created"

	usecase := bootstrapp_sa2.NewBootstrapSuperAdmin(userGateway, idGenerator, encoder)
	input := bootstrapp_sa2.BootstrapSAInput{
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

	if output == nil {
		t.Errorf("output should not be nil")
	}

	if !userGateway.ExistsActiveSuperAdminResult {
		t.Error("ExistsActiveSuperAdmin should return true")
	}

	if !userGateway.ExistsActiveSuperAdminCalled {
		t.Error("ExistsActiveSuperAdmin should have been called")
	}

	if output.Message != expectedOutput {
		t.Errorf("Should return error message: %v", output.Message)
	}

}

func TestShouldReturnErrorIfIdGeneratorFails(t *testing.T) {
	//Arrange
	expectedError := shared2.InternalError(shared2.ErrEmptyID)
	username := "superAdmin"
	userGateway := &mocks.UserGatewayMock{}
	idGenerator := &mocks.IDGeneratorMock{}
	encoder := &mocks.EncoderGatewayMock{}
	usecase := bootstrapp_sa2.NewBootstrapSuperAdmin(userGateway, idGenerator, encoder)
	input := bootstrapp_sa2.BootstrapSAInput{
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
	userGateway := &mocks.UserGatewayMock{}
	idGenerator := &mocks.IDGeneratorMock{GenerateResult: "test-id"}
	encodeErr := errors.New("test-error")
	expectedError := shared2.InternalError(encodeErr)
	encodeParam := "str0ngP@ssword"
	encoder := &mocks.EncoderGatewayMock{
		EncodeParam: encodeParam,
		EncodeErr:   encodeErr,
	}

	usecase := bootstrapp_sa2.NewBootstrapSuperAdmin(userGateway, idGenerator, encoder)
	input := bootstrapp_sa2.BootstrapSAInput{
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
	userGateway := &mocks.UserGatewayMock{}
	idGenerator := &mocks.IDGeneratorMock{GenerateResult: "test-id"}
	encodeParam := "str0ngP@ssword"
	encodedPassword := "encodedPassword"
	encoder := &mocks.EncoderGatewayMock{EncodeResult: encodedPassword}
	expectedError := shared2.RequiredField("User", "name")

	usecase := bootstrapp_sa2.NewBootstrapSuperAdmin(userGateway, idGenerator, encoder)
	input := bootstrapp_sa2.BootstrapSAInput{
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
	userGateway := &mocks.UserGatewayMock{}
	idGenerator := &mocks.IDGeneratorMock{GenerateResult: "test-id"}
	encodeParam := "str0ngP@ssword"
	encodedPassword := "encodedPassword"
	encoder := &mocks.EncoderGatewayMock{EncodeResult: encodedPassword}
	expectedError := shared2.RequiredField("User", "email")

	usecase := bootstrapp_sa2.NewBootstrapSuperAdmin(userGateway, idGenerator, encoder)
	input := bootstrapp_sa2.BootstrapSAInput{
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
	userGateway := &mocks.UserGatewayMock{}
	idGenerator := &mocks.IDGeneratorMock{GenerateResult: "test-id"}
	encodeParam := "str0ngP@ssword"
	encodedPassword := "encodedPassword"
	encoder := &mocks.EncoderGatewayMock{EncodeResult: encodedPassword}
	expectedError := shared2.RequiredField("User", "username")

	usecase := bootstrapp_sa2.NewBootstrapSuperAdmin(userGateway, idGenerator, encoder)
	input := bootstrapp_sa2.BootstrapSAInput{
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

func TestShouldReturnErrorIfPasswordIsEmpty(t *testing.T) {
	//Arrange
	username := ""
	userGateway := &mocks.UserGatewayMock{}
	idGenerator := &mocks.IDGeneratorMock{GenerateResult: "test-id"}
	encodeParam := "str0ngP@ssword"
	encodedPassword := "encodedPassword"
	encoder := &mocks.EncoderGatewayMock{EncodeResult: encodedPassword}
	expectedError := shared2.RequiredField("User", "username")

	usecase := bootstrapp_sa2.NewBootstrapSuperAdmin(userGateway, idGenerator, encoder)
	input := bootstrapp_sa2.BootstrapSAInput{
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

func TestShouldReturnErrorIfSaveFails(t *testing.T) {
	//Arrange
	username := "any_username"
	errr := errors.New("some error")
	expectedError := shared2.InternalError(errr)
	userGateway := &mocks.UserGatewayMock{SaveError: errr}
	userId := "test-id"
	idGenerator := &mocks.IDGeneratorMock{GenerateResult: userId}
	encodeParam := "str0ngP@ssword"
	encodedPassword := "encodedPassword"
	encoder := &mocks.EncoderGatewayMock{EncodeResult: encodedPassword}

	usecase := bootstrapp_sa2.NewBootstrapSuperAdmin(userGateway, idGenerator, encoder)
	input := bootstrapp_sa2.BootstrapSAInput{
		Username: username,
		Name:     "System Admin",
		Email:    "sa@system.com",
		Password: encodeParam,
	}

	//Act
	output, err := usecase.Execute(input)

	//Assert
	if err == nil {
		t.Errorf("should return an error")
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

	if idGenerator.GenerateResult != userId {
		t.Errorf("Generate should return %v", userId)
	}

	if !encoder.EncodeCalled {
		t.Error("Encoder should have been called")
	}

	if encoder.EncodeResult != encodedPassword {
		t.Errorf("Encoder should return %v", encodedPassword)
	}

	if !userGateway.SaveCalled {
		t.Error("Save should have been called")
	}

	if err.Code != expectedError.Code {
		t.Errorf("Should return error code: %v", expectedError.Code)
	}

	if err.Message != expectedError.Message {
		t.Errorf("Should return error message: %v", expectedError.Message)
	}
}

func TestShouldSaveSa(t *testing.T) {
	//Arrange
	username := "any_username"
	userGateway := &mocks.UserGatewayMock{}
	userId := "test-id"
	idGenerator := &mocks.IDGeneratorMock{GenerateResult: userId}
	encodeParam := "str0ngP@ssword"
	encodedPassword := "encodedPassword"
	encoder := &mocks.EncoderGatewayMock{EncodeResult: encodedPassword}
	expectedOutput := "Super Admin successfully created"

	usecase := bootstrapp_sa2.NewBootstrapSuperAdmin(userGateway, idGenerator, encoder)
	input := bootstrapp_sa2.BootstrapSAInput{
		Username: username,
		Name:     "System Admin",
		Email:    "sa@system.com",
		Password: encodeParam,
	}

	//Act
	output, err := usecase.Execute(input)

	//Assert
	if err != nil {
		t.Errorf("should not return an error")
	}

	if output == nil {
		t.Errorf("output should not be nil")
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

	if idGenerator.GenerateResult != userId {
		t.Errorf("Generate should return %v", userId)
	}

	if !encoder.EncodeCalled {
		t.Error("Encoder should have been called")
	}

	if encoder.EncodeResult != encodedPassword {
		t.Errorf("Encoder should return %v", encodedPassword)
	}

	if !userGateway.SaveCalled {
		t.Error("Save should have been called")
	}

	if output.Message != expectedOutput {
		t.Errorf("Should return error code: %v", output.Message)
	}
}

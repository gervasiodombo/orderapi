package user_test

import (
	"testing"

	"github.com/oderapi/domain/entity/user"
	"github.com/oderapi/domain/shared"
)

//NewFirstSuperAdmin testes

func TestShouldReturnDomainErrorIfNameIsEmpty(t *testing.T) {
	//Arrange
	expectedCode := shared.BadRequest
	expectedMessage := `The field 'name' is required in 'User'`

	//Act
	id, _ := shared.NewID("test-sa-id")
	_, err := user.NewFirstSuperAdmin(id, "", "any_email", "any_password")

	//Assert
	if err == nil {
		t.Error("expected error, got none")
	}
	if err.Code != expectedCode {
		t.Errorf("expected BadRequest, got %s", err.Code)
	}
	if err.Message != expectedMessage {
		t.Errorf("expected %s, got %s", expectedCode, err.Message)
	}
}

func TestShouldReturnDomainErrorIfEmailIsEmpty(t *testing.T) {
	//Arrange
	expectedCode := shared.BadRequest
	expectedMessage := `The field 'email' is required in 'User'`

	//Atc
	id, _ := shared.NewID("test-sa-id")
	_, err := user.NewFirstSuperAdmin(id, "any_name", "", "any_password")

	//Assert
	if err == nil {
		t.Error("expected error, got none")
	}
	if err.Code != expectedCode {
		t.Errorf("expected BadRequest, got %s", err.Code)
	}
	if err.Message != expectedMessage {
		t.Errorf("expected %s, got %s", expectedCode, err.Message)
	}
}

func TestShouldReturnDomainErrorIfPasswordIsEmpty(t *testing.T) {
	//Arrange
	expectedCode := shared.BadRequest
	expectedMessage := `The field 'password' is required in 'User'`

	//Atc
	id, _ := shared.NewID("test-sa-id")
	_, err := user.NewFirstSuperAdmin(id, "any_name", "any_email", "")

	//Assert
	if err == nil {
		t.Error("expected error, got none")
	}
	if err.Code != expectedCode {
		t.Errorf("expected BadRequest, got %s", err.Code)
	}
	if err.Message != expectedMessage {
		t.Errorf("expected %s, got %s", expectedCode, err.Message)
	}
}

package user_test

import (
	"testing"

	"github.com/oderapi/domain/entity/user"
	"github.com/oderapi/domain/shared"
)

func TestShouldReturnDomainErrorIfNameIsEmpty(t *testing.T) {
	//Arrange
	expectedCode := shared.BadRequest
	expectedMessage := `The field 'name' is required in 'User'`

	//Act
	id, _ := shared.NewID("test-sa-id")
	_, err := user.NewFirstSuperAdmin(id, "", "any_email", "any_username", "any_password")

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
	_, err := user.NewFirstSuperAdmin(id, "any_name", "", "any_username", "any_password")

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

func TestShouldReturnDomainErrorIfUsernamesEmpty(t *testing.T) {
	//Arrange
	expectedCode := shared.BadRequest
	expectedMessage := `The field 'username' is required in 'User'`

	//Atc
	id, _ := shared.NewID("test-sa-id")
	_, err := user.NewFirstSuperAdmin(id, "any_name", "any_email", "", "any_password")

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
	_, err := user.NewFirstSuperAdmin(id, "any_name", "any_email", "any_username", "")

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

func TestShouldReturnUserWithCorrectValues(t *testing.T) {
	//Atc 	//Arrange
	id, _ := shared.NewID("test-sa-id")
	user, err := user.NewFirstSuperAdmin(id, "any_name", "any_email", "any_username", "any_passq")

	//Assert
	if err != nil {
		t.Error("should not return error")
	}

	if user.Id().Value() != id.Value() {
		t.Errorf("expected : %s, but got %s", id, user.Id())
	}

	if user.Name() != "any_name" {
		t.Errorf("expected : %s, but got %s", "any_name", user.Name())
	}

	if user.Email() != "any_email" {
		t.Errorf("expected : %s, but got %s", "any_email", user.Email())
	}

	if user.Username() != "any_username" {
		t.Errorf("expected : %s, but got %s", "any_username", user.Username())
	}

	if user.Password() != "any_passq" {
		t.Errorf("expected : %s, but got %s", "any_passq", user.Password())
	}

	if user.Status().String() != "ACTIVE" {
		t.Errorf("expected : %s, but got %s", "ACTIVE", user.Status().String())
	}

	if len(user.Roles()) != 1 {
		t.Errorf("expected : %d, but got %d", 1, len(user.Roles()))
	}

	if user.Roles()[0].String() != "SUPER_ADMIN" {
		t.Errorf("expected : %s, but got %s", "SUPER_ADMIN", user.Roles()[0].String())
	}
}

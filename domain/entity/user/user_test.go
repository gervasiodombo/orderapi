package user_test

import (
	"testing"

	"github.com/oderapi/domain/entity/user"
	"github.com/oderapi/domain/shared"
)

// NewFirstSuperAdmin
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
	us, err := user.NewFirstSuperAdmin(id, "any_name", "any_email", "any_username", "any_passq")

	//Assert
	if err != nil {
		t.Error("should not return error")
	}

	if us.Id().Value() != id.Value() {
		t.Errorf("expected : %s, but got %s", id, us.Id())
	}

	if us.Name() != "any_name" {
		t.Errorf("expected : %s, but got %s", "any_name", us.Name())
	}

	if us.Email() != "any_email" {
		t.Errorf("expected : %s, but got %s", "any_email", us.Email())
	}

	if us.Username() != "any_username" {
		t.Errorf("expected : %s, but got %s", "any_username", us.Username())
	}

	if us.Password() != "any_passq" {
		t.Errorf("expected : %s, but got %s", "any_passq", us.Password())
	}

	if us.Status().String() != "ACTIVE" {
		t.Errorf("expected : %s, but got %s", "ACTIVE", us.Status().String())
	}

	if len(us.Roles()) != 1 {
		t.Errorf("expected : %d, but got %d", 1, len(us.Roles()))
	}

	if us.Roles()[0].String() != "SUPER_ADMIN" {
		t.Errorf("expected : %s, but got %s", "SUPER_ADMIN", us.Roles()[0].String())
	}
}

// With
func TestShouldReturnUserWithIncorrectValues(t *testing.T) {
	id, _ := shared.NewID("test-sa-id")
	us := user.With(id, "any_name", "any_email", "any_username", "any_passq", user.ACTIVE, []user.Role{user.SUPER_ADMIN})

	if us.Id().Value() != id.Value() {
		t.Errorf("expected : %s, but got %s", id, us.Id())
	}

	if us.Name() != "any_name" {
		t.Errorf("expected : %s, but got %s", "any_name", us.Name())
	}

	if us.Email() != "any_email" {
		t.Errorf("expected : %s, but got %s", "any_email", us.Email())
	}

	if us.Username() != "any_username" {
		t.Errorf("expected : %s, but got %s", "any_username", us.Username())
	}

	if us.Password() != "any_passq" {
		t.Errorf("expected : %s, but got %s", "any_passq", us.Password())
	}

	if us.Status().String() != "ACTIVE" {
		t.Errorf("expected : %s, but got %s", "ACTIVE", us.Status().String())
	}

	if len(us.Roles()) != 1 {
		t.Errorf("expected : %d, but got %d", 1, len(us.Roles()))
	}

	if us.Roles()[0].String() != "SUPER_ADMIN" {
		t.Errorf("expected : %s, but got %s", "SUPER_ADMIN", us.Roles()[0].String())
	}
}

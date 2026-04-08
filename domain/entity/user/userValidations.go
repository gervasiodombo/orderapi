package user

import "github.com/oderapi/domain/shared"

func validateName(name string) (string, *shared.DomainError) {
	if name == "" {
		return "", shared.RequiredField("User", "name")
	}
	return name, nil
}

func validateEmail(email string) (string, *shared.DomainError) {
	if email == "" {
		return "", shared.RequiredField("User", "email")
	}
	return email, nil
}

func validateUsername(username string) (string, *shared.DomainError) {
	if username == "" {
		return "", shared.RequiredField("User", "username")
	}
	return username, nil
}

func validatePassword(password string) (string, *shared.DomainError) {
	if password == "" {
		return "", shared.RequiredField("User", "password")
	}
	return password, nil
}

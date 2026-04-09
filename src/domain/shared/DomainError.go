package shared

import (
	"fmt"
)

type DomainError struct {
	Code    ErrorCode
	Message string
	Cause   error
}

func RequiredField(entity string, name string) *DomainError {
	return &DomainError{
		BadRequest,
		fmt.Sprintf("The field '%s' is required in '%s'", name, entity),
		fmt.Errorf("the field '%s' is required in '%s'", name, entity),
	}
}

func InternalError(err error) *DomainError {
	return &DomainError{
		InterServerError,
		fmt.Sprintf("An unexpected error occurred while processing your request, Please try again later if the problem persists please contact the system administrator."),
		err,
	}
}

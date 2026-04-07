package shared

type ErrorCode string

const (
	BadRequest       ErrorCode = "BadRequest"
	InterServerError ErrorCode = "InterServerError"
)

func (code ErrorCode) HTTPStatus() int {
	switch code {
	case BadRequest:
		return 400
	default:
		return 500
	}
}

func (code ErrorCode) String() string {
	return string(code)
}

package user

type Status string

const (
	ACTIVE   Status = "ACTIVE"
	INACTIVE Status = "INACTIVE"
)

func (s Status) String() string {
	return string(s)
}

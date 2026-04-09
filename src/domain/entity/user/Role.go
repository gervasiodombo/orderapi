package user

type Role string

const (
	ADMIN       Role = "ADMIN"
	SUPER_ADMIN Role = "SUPER_ADMIN"
	CUSTOMER    Role = "CUSTOMER"
)

func (r Role) String() string {
	return string(r)
}

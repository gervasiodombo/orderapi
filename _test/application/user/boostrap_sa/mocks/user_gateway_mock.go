package mocks

import (
	"github.com/oderapi/domain/entity/user"
)

type UserGatewayMock struct {
	ExistsActiveSuperAdminResult bool
	ExistsActiveSuperAdminCalled bool
	ExistsActiveSuperAdminParam  string
	SaveParam                    user.User
	SaveError                    error
	SaveCalled                   bool
}

func (m *UserGatewayMock) ExistsActiveSuperAdmin(username string) bool {
	m.ExistsActiveSuperAdminCalled = true
	m.ExistsActiveSuperAdminParam = username
	return m.ExistsActiveSuperAdminResult
}

func (m *UserGatewayMock) Save(user user.User) error {
	m.SaveCalled = true
	m.SaveParam = user
	return m.SaveError
}

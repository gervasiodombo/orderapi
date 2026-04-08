package _mocks

import (
	"github.com/oderapi/domain/entity/user"
)

type UserGatewayMock struct {
	ExistsActiveSuperAdminResult bool
	ExistsActiveSuperAdminCalled bool
	ExistsActiveSuperAdminErr    error
	SaveParam                    user.User
	SaveError                    error
	SaveCalled                   bool
}

func (m *UserGatewayMock) ExistsActiveSuperAdmin() (bool, error) {
	m.ExistsActiveSuperAdminCalled = true
	return m.ExistsActiveSuperAdminResult, m.ExistsActiveSuperAdminErr
}

func (m *UserGatewayMock) Save(user user.User) error {
	m.SaveCalled = true
	m.SaveParam = user
	return m.SaveError
}

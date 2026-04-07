package mocks

import (
	"github.com/oderapi/domain/entity/user"
	"github.com/oderapi/domain/vo"
)

type UserGatewayMock struct {
	ExistsActiveSuperAdminResult bool
	ExistsActiveSuperAdminCalled bool
	ExistsActiveSuperAdminParam  string
	SaveResult                   *vo.Output
	SaveParam                    user.User
	SaveError                    error
	SaveCalled                   bool
}

func (m *UserGatewayMock) ExistsActiveSuperAdmin(username string) bool {
	m.ExistsActiveSuperAdminCalled = true
	m.ExistsActiveSuperAdminParam = username
	return m.ExistsActiveSuperAdminResult
}

func (m *UserGatewayMock) Save(user user.User) (*vo.Output, error) {
	m.SaveCalled = true
	m.SaveParam = user
	return m.SaveResult, m.SaveError
}

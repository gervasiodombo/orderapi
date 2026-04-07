package mocks

import (
	"github.com/oderapi/domain/entity/user"
	"github.com/oderapi/domain/shared"
)

type UserGatewayMock struct {
	ExistsActiveSuperAdminResult bool
	ExistsActiveSuperAdminCalled bool
	ExistsActiveSuperAdminParam  string

	// Save
	SaveCalled bool
	SavedUser  user.User
}

func (m *UserGatewayMock) ExistsActiveSuperAdmin(username string) bool {
	m.ExistsActiveSuperAdminCalled = true
	m.ExistsActiveSuperAdminParam = username
	return m.ExistsActiveSuperAdminResult
}

func (m *UserGatewayMock) Save(u user.User) {
	m.SaveCalled = true
	m.SavedUser = u
}

func (m *UserGatewayMock) FindByEmail(email string) (*user.User, *shared.DomainError) {
	return nil, nil
}

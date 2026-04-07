package mocks

type UserGatewayMock struct {
	ExistsActiveSuperAdminResult bool
	ExistsActiveSuperAdminCalled bool
	ExistsActiveSuperAdminParam  string
}

func (m *UserGatewayMock) ExistsActiveSuperAdmin(username string) bool {
	m.ExistsActiveSuperAdminCalled = true
	m.ExistsActiveSuperAdminParam = username
	return m.ExistsActiveSuperAdminResult
}

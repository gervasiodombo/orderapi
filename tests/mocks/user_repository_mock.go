package mocks

import (
	"github.com/oderapi/src/domain/entity/user"
)

type UserRepositoryMock struct {
	ExistsResult bool
	ExistsErr    error
	SaveErr      error
}

func (m *UserRepositoryMock) ExistsActiveSuperAdmin() (bool, error) {
	return m.ExistsResult, m.ExistsErr
}

func (m *UserRepositoryMock) Save(u user.User) error {
	return m.SaveErr
}

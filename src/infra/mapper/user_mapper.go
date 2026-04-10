package mapper

import (
	"github.com/oderapi/src/domain/entity/user"
	"github.com/oderapi/src/domain/shared"
	"github.com/oderapi/src/infra/persistence/model"
)

func ToModel(u user.User) model.UserModel {
	roles := make([]model.RoleModel, len(u.Roles()))
	for i, role := range u.Roles() {
		roles[i] = model.RoleModel{
			UserID: u.Id().Value(),
			Role:   role.String(),
		}
	}

	return model.UserModel{
		ID:       u.Id().Value(),
		Name:     u.Name(),
		Email:    u.Email(),
		Username: u.Username(),
		Password: u.Password(),
		Status:   u.Status().String(),
		Roles:    roles,
	}
}

func ToDomain(m model.UserModel) (user.User, error) {
	id, err := shared.NewID(m.ID)
	if err != nil {
		return user.User{}, err
	}

	roles := make([]user.Role, len(m.Roles))
	for i, r := range m.Roles {
		roles[i] = user.Role(r.Role)
	}

	return user.With(id, m.Name, m.Email, m.Username, m.Password, user.Status(m.Status), roles), nil
}

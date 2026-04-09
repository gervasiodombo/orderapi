package mapper

import (
	user2 "github.com/oderapi/src/domain/entity/user"
	"github.com/oderapi/src/domain/shared"
	model2 "github.com/oderapi/src/infra/persistence/model"
)

func ToModel(u user2.User) model2.UserModel {
	roles := make([]model2.RoleModel, len(u.Roles()))
	for i, role := range u.Roles() {
		roles[i] = model2.RoleModel{
			UserID: u.Id().Value(),
			Role:   role.String(),
		}
	}

	return model2.UserModel{
		ID:       u.Id().Value(),
		Name:     u.Name(),
		Email:    u.Email(),
		Username: u.Username(),
		Password: u.Password(),
		Status:   u.Status().String(),
		Roles:    roles,
	}
}

func ToDomain(m model2.UserModel) (user2.User, error) {
	id, err := shared.NewID(m.ID)
	if err != nil {
		return user2.User{}, err
	}

	roles := make([]user2.Role, len(m.Roles))
	for i, r := range m.Roles {
		roles[i] = user2.Role(r.Role)
	}

	return user2.With(id, m.Name, m.Email, m.Username, m.Password, user2.Status(m.Status), roles), nil
}

package repository

import (
	"fmt"

	"github.com/oderapi/src/domain/entity/user"
	"github.com/oderapi/src/infra/mapper"
	"github.com/oderapi/src/infra/persistence/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) ExistsActiveSuperAdmin() (bool, error) {
	var count int64
	result := r.db.Model(&model.UserModel{}).
		Joins(`JOIN "T_USER_ROLES" ur ON ur.user_id = "T_USERS".id`).
		Where(`"T_USERS".status = ? AND ur.role = ?`, "ACTIVE", "SUPER_ADMIN").
		Count(&count)
	if result.Error != nil {
		return false, fmt.Errorf("failed to check active super admin: %w", result.Error)
	}
	return count > 0, nil
}

func (r *UserRepositoryImpl) Save(u user.User) error {
	userModel := mapper.ToModel(u)
	result := r.db.Create(userModel)
	if result.Error != nil {
		return fmt.Errorf("failed to save user: %w", result.Error)
	}
	return nil
}

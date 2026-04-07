package repository

import (
	"fmt"

	"github.com/oderapi/domain/entity/user"
	"github.com/oderapi/infra/mapper"
	"github.com/oderapi/infra/persistence/model"
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
		Joins("JOIN user_roles ur ON ur.user_id = users.id").
		Where("users.status = ? AND ur.role = ?", "ACTIVE", "SUPER_ADMIN").
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

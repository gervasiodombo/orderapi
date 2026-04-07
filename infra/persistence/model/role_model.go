package model

type RoleModel struct {
	UserID string `gorm:"column:user_id;primaryKey"`
	Role   string `gorm:"column:role;primaryKey"`
}

func (RoleModel) TableName() string {
	return "T_USER_ROLES"
}

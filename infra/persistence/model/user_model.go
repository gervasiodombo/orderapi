package model

type UserModel struct {
	ID       string      `gorm:"primaryKey;column:id"`
	Name     string      `gorm:"column:name;not null"`
	Email    string      `gorm:"column:email;not null;unique"`
	Username string      `gorm:"column:username;not null;unique"`
	Password string      `gorm:"column:password;not null"`
	Status   string      `gorm:"column:status;not null;default:ACTIVE"`
	Roles    []RoleModel `gorm:"foreignKey:UserID"`
}

func (UserModel) TableName() string {
	return "T_USERS"
}

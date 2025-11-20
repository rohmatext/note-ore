package entity

type Role struct {
	ID    uint   `gorm:"column:id;primaryKey"`
	Name  string `gorm:"column:name"`
	Users []User `gorm:"foreignKey:RoleID"`
}

func (Role) TableName() string {
	return "roles"
}

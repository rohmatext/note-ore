package seeder

import (
	"fmt"
	"rohmatext/ore-note/internal/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func RoleSeeder(db *gorm.DB) {
	var roles = []entity.Role{{Name: "admin"}, {Name: "operator"}}

	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoNothing: true,
	}).Create(&roles).Error; err != nil {
		panic(fmt.Sprintf("RoleSeeder: failed to create roles: %v", err))
	}

	fmt.Printf("RoleSeeder: successfully seeded %d roles\n", len(roles))
}

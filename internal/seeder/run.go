package seeder

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	fmt.Println("Running seeders...")
	db.WithContext(context.Background()).Transaction(func(tx *gorm.DB) error {
		RoleSeeder(tx)
		UserSeeder(tx)
		return nil
	})
	fmt.Println("Done seeding!")
}

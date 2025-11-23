package seeder

import (
	"fmt"
	"math/rand"
	"rohmatext/ore-note/internal/entity"
	"strings"

	"github.com/go-faker/faker/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func fakeName() string {
	firstName := faker.FirstName()
	lastName := []string{faker.LastName(), ""}
	fullName := fmt.Sprintf("%s %s", firstName, lastName[rand.Intn(len(lastName))])
	return strings.Trim(fullName, " ")
}

func UserSeeder(db *gorm.DB) {
	var roles []entity.Role
	if err := db.Find(&roles).Error; err != nil {
		panic(fmt.Sprintf("UserSeeder: failed to read roles from database: %+v\n", err))
	}

	password, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		panic(fmt.Sprintf("UserSeeder: failed to generate password: %+v", err))
	}

	users := make([]entity.User, len(roles))
	for idx, role := range roles {
		users[idx] = entity.User{
			Name:     fakeName(),
			Username: role.Name,
			Password: string(password),
			Role:     role,
		}
	}

	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "username"}},
		DoNothing: true,
	}).Create(&users).Error; err != nil {
		panic(fmt.Sprintf("UserSeeder: failed to create users: %v", err))
	}

	fmt.Printf("UserSeeder: successfully seeded %d users\n", len(users))
}

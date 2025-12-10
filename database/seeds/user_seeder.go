package seeds

import (
	"bwanews/internal/core/domain/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/rs/zerolog/log"
)

func SeedRoles(db *gorm.DB) {
	// Implementation for seeding roles

	bytes, err := bcrypt.GenerateFromPassword([]byte("admin123"), 14)
	
	if err != nil {
		log.Fatal().Err(err).Msg("[UserSeeder] Failed to hash password")
		return
	}

	admin := []model.User{
		{
			Name:     "Admin",
			Email:    "admin@mail.com",
			Password: string(bytes),
		},
	}

	if err := db.FirstOrCreate(&admin, model.User{Email: "admin@mail.com"}).Error; err != nil {
		log.Fatal().Err(err).Msg("[UserSeeder] Failed to seed admin user")	
	} else {
		log.Info().Msg("[UserSeeder] Successfully seeded admin user")
	}
}
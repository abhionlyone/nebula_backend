package migrations

import (
	"nebula_backend/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	m := &gormigrate.Migration{
		ID: "20240613145315",
		Migrate: func(tx *gorm.DB) error {
			// Add the 'age' column to the 'users' table
			return tx.Migrator().AddColumn(&models.User{}, "Age")
		},
		Rollback: func(tx *gorm.DB) error {
			// Remove the 'age' column from the 'users' table
			return tx.Migrator().DropColumn(&models.User{}, "Age")
		},
	}

	Migrations = append(Migrations, m)
}

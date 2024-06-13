// package migrations

// import (
// 	"log"
// 	"fmt"
// 	"nebula_backend/config"
// 	"nebula_backend/models"
// 	"github.com/go-gormigrate/gormigrate/v2"
// 	"gorm.io/gorm"
// )

// func Migrate() {
// 	m := gormigrate.New(config.DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
// 		{
// 			ID: "2023061301",
// 			Migrate: func(tx *gorm.DB) error {
// 				return tx.AutoMigrate(&models.User{})
// 			},
// 			Rollback: func(tx *gorm.DB) error {
// 				return tx.Migrator().DropTable("users")
// 			},
// 		},
// 	})

// 	if err := m.Migrate(); err != nil {
// 		log.Fatalf("Could not migrate: %v", err)
// 	}
// 	fmt.Println("Migration did run successfully")
// }

package migrations

import (
	"fmt"
	"log"
	"nebula_backend/config"
	"nebula_backend/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var Migrations []*gormigrate.Migration

func init() {
	Migrations = []*gormigrate.Migration{
		{
			ID: "2023061301",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
	}
}

func Migrate() {
	m := gormigrate.New(config.DB, gormigrate.DefaultOptions, Migrations)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	fmt.Println("Migration did run successfully")
}


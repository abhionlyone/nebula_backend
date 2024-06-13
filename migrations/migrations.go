package migrations

import (
	"fmt"
	"log"
	"nebula_backend/config"

	"github.com/go-gormigrate/gormigrate/v2"
)

var Migrations []*gormigrate.Migration

func init() {
	// Migrations will be appended from individual migration files
}

func Migrate() {
	m := gormigrate.New(config.DB, gormigrate.DefaultOptions, Migrations)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	fmt.Println("Migration did run successfully")
}
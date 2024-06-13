package main

import (
	"log"
	"nebula_backend/config"
	"nebula_backend/migrations"
	"nebula_backend/routes"
	"nebula_backend/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	utils.InitLogger()
	config.InitDB()
	migrations.Migrate()

	r := routes.SetupRouter()
	r.Run(":8080")
}

package main

import (
	"nebula_backend/config"
	"nebula_backend/migrations"
	"nebula_backend/routes"
	"nebula_backend/utils"
)

func main() {
	utils.InitLogger()
	config.InitDB()
	migrations.Migrate()

	r := routes.SetupRouter()
	r.Run(":8080")
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"simkes-go/config/database"
	"simkes-go/migration"
	"simkes-go/routes"
)

func main() {
	env := godotenv.Load()
	database.ConnectDB()

	migrate := os.Getenv("MIGRATE")
	if migrate == "TRUE" {
		migration.RunMigration()
	}
	log.Println("Run System")
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})
	routes.RouteInit(app)

	if env != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}

	listen := app.Listen(":" + port)
	if listen != nil {
		log.Println("Fail to listen go fiber server")
		os.Exit(1000)
	}
}

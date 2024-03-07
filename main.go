package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/quocquann/book-api/pkg/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Fail to load env")
	}
	app := fiber.New()

	routes.PublicRoute(app)

	app.Listen(":8080")
}

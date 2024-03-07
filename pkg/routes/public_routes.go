package routes

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/quocquann/book-api/app/handlers"
)

func PublicRoute(app *fiber.App) {
	v1 := app.Group("/api/v1")
	{
		v1.Get("/books", handler.GetBooks)
		v1.Get("/books/:isbn", handler.GetBookByIsbn)
	}

}

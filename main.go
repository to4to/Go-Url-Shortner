package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/to4to/go-url-shortner/api/routes"
)



func setupRoutes(app *fiber.App) {
	app.Get("/:url",routes.ResolveURL)
	app.Get("/api/v1",rouroutes.ShortenURL)
}



func main(){

godotenv.Load()


app:=fiber.New()

}
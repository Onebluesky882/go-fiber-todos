package posts

import "github.com/gofiber/fiber/v2"

func RegisterPostRouted(app *fiber.App) {
	PostController := NewPostController()
	route := app.Group("/posts")
	route.Get("/", PostController.GetPosts)
}

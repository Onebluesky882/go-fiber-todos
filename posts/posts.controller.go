package posts

import "github.com/gofiber/fiber/v2"

type PostController struct{}

func NewPostController() *PostController {
	return &PostController{}
}

func (ctl *PostController) GetPosts(c *fiber.Ctx) error {
	data, err := NewServicePost()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch posts",
		})
	}
	return c.SendString(data)
}

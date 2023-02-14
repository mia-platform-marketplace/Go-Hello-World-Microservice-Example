package main

import (
	swagger "github.com/davidebianchi/gswagger"
	"github.com/gofiber/fiber/v2"
)

func HelloHandler(c *fiber.Ctx) error {
	return c.JSON(HelloWorld{"Hello world!"})
}

var HelloSchema = swagger.Definitions{
	Responses: map[int]swagger.ContentValue{
		200: {
			Content: swagger.Content{
				"application/json": {Value: &HelloWorld{}},
			},
		},
	},
}

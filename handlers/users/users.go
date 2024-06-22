package handlers

import (
	// "fmt"
	// "net/http"
	user "karma_auth/helpers/users"

	"github.com/gofiber/fiber/v2"
)

// ResponseHTTP represents response body of this API
type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func CreateUser(c *fiber.Ctx) error {
	users, err := user.GetUsers()
	if err == nil {
		return c.JSON(ResponseHTTP{
			Success: true,
			Message: "Successfully retrieved all users.",
			Data:    users,
		})
	} else {
		return c.JSON(ResponseHTTP{
			Success: true,
			Message: "Failed to retrieve all users.",
			Data:    nil,
		})
	}
}
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

type CreateUserRequest struct {
	Name	 string `json:"name" form:"name"`
	Email	 string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone	 string `json:"phone" form:"phone"`
	Details	 string `json:"details" form:"details"`
}

func GetUsers(c *fiber.Ctx) error {
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

func CreateUser(c *fiber.Ctx) error {
	req := new(CreateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(ResponseHTTP{
			Success: false,
			Message: "Failed to parse request body.",
			Data:    nil,
		})
	}
	oid := c.Locals("oid").(string)
	user.CreateUser(oid, req.Name, req.Email, req.Password, req.Phone, req.Details)
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully created user.",
		Data:    nil,
	})
}
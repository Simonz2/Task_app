package handlers

import (
	"github.com/Simonz2/Task_app/pkg/config"
	"github.com/Simonz2/Task_app/pkg/repo"
	"github.com/Simonz2/Task_app/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *fiber.Ctx) error {
	//handle userlogin
	var req request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	//access the database
	userRepo := repo.NewUserRepo(config.GetDB())

	//check if user exists
	user, err := userRepo.GetUserByUsername(req.Username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "username does not exist"})
	}
	//check if user is active
	if !user.Activated {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "user is not active"})
	}
	//check if password is correct
	if !user.CheckPassword(req.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid password"})
	}

	//create a token for the user
	token, err := utils.CreateToken(user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token, "message": "login successful"})
}

//password with min: 8 chars, at least one uppercase, one digit, one special char

func RegisterHandler(c *fiber.Ctx) error {
	var req request

	//check if the request body is valid
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request payload"})
	}
	//basic field validation
	if req.Username == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "username or password is empty"})
	}
	//
	userRepo := repo.NewUserRepo(config.GetDB())
	err := userRepo.CreateUser(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "user created successfully"})
}

package handlers

import (
	"github.com/Simonz2/Task_app/pkg/config"
	"github.com/Simonz2/Task_app/pkg/repo"
	"github.com/Simonz2/Task_app/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	//handle userlogin
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
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

func RegisterHandler(c *fiber.Ctx) error {
	return nil
}

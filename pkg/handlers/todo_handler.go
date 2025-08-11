package handlers

import (
	"strconv"

	"github.com/Simonz2/Task_app/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// get all todos
func GetTodos(c *fiber.Ctx) error {
	todos := models.GetTodos()
	return c.Status(200).JSON(todos)
}

// create a new todo
func CreateTodo(c *fiber.Ctx) error {
	todo := &models.Todo{}
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	if todo.Body == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Todo body is required"})
	}
	todo.CreateTodo()
	return c.Status(fiber.StatusCreated).JSON(todo)
}

// update a todo by id(todo.completed=true)
func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todo := &models.Todo{}
	ids, _ := strconv.ParseInt(id, 10, 64)
	todo = models.PatchTodo(ids)
	if todo != nil {
		return c.Status(200).JSON(&todo)
	}
	return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
}

// delete a todo by id
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	ids, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if models.DeleteTodo(ids) != (models.Todo{}) {
		return c.Status(200).JSON(fiber.Map{"success": true})
	}
	return c.Status(404).JSON(fiber.Map{"error": "todo not found"})

}

package routes

import (
	"strconv"

	"github.com/Simonz2/Task_app/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// In-memory todos slice (for demonstration; consider moving to a service or DB layer)
var todos = []models.Todo{}

// RegisterTodoRoutes sets up all todo-related routes
func RegisterTodoRoutes(app *fiber.App) {
	//get all todos not deleted
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		todos = models.GetTodos()
		return c.Status(200).JSON(todos)
	})
	//create a new todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &models.Todo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}
		todo.CreateTodo()
		return c.Status(201).JSON(todo)
	})
	//update a todo by id(todo.completed=true)
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		todo := &models.Todo{}
		ids, _ := strconv.ParseInt(id, 10, 64)
		todo = models.PatchTodo(ids)
		if todo != nil {
			return c.Status(200).JSON(&todo)
		}
		return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
	})
	//delete a todo by id
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		ids, _ := strconv.ParseInt(id, 10, 64)
		if models.DeleteTodo(ids) != (models.Todo{}) {
			return c.Status(200).JSON(fiber.Map{"success": true})
		}
		return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
	})
}

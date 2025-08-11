package routes

import (
	"github.com/Simonz2/Task_app/pkg/handlers"
	"github.com/Simonz2/Task_app/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

// RegisterTodoRoutes sets up all todo-related routes
func RegisterTodoRoutes(app *fiber.App) {
	api := app.Group("/api")                                  // Apply authentication middleware
	api.Post("/login", handlers.LoginHandler)                 // Login route
	api.Post("/register", handlers.RegisterHandler)           // Register route
	protected := api.Group("/todos", middleware.Authenticate) // Protected routes for todos
	protected.Get("/", handlers.GetTodos)                     // Get all todos
	protected.Post("/", handlers.CreateTodo)                  // Create a new todo             // Get todo by ID
	protected.Patch("/:id", handlers.UpdateTodo)              // Update todo by ID
	protected.Delete("/:id", handlers.DeleteTodo)             // Delete todo by ID
}

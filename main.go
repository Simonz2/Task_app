package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Simonz2/Task_app/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("now connected to DB todos")
	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	routes.RegisterTodoRoutes(app)
	log.Fatal(app.Listen(":" + PORT))
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Simonz2/Task_app/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("now connected to DB todos")
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	routes.RegisterTodoRoutes(app)
	log.Fatal(app.Listen(":" + PORT))
}

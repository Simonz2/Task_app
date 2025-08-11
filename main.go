package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Simonz2/Task_app/pkg/routes"
	"github.com/Simonz2/Task_app/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("now connected to DB todos")
	utils.InitJWT() //initialize JWT secret key
	app := fiber.New()
	/*
	 */
	if os.Getenv("ENV") != "production" {
		//load the .env file if not in production
		_ = godotenv.Load(".env")
		//only when developing else the cors will not work
		app.Use(cors.New(cors.Config{
			AllowOrigins: "http://localhost:5173",
			AllowHeaders: "Origin, Content-Type, Accept",
		}))
		/*
			if err != nil {
				log.Fatal("Error loading .env file")
			}
		*/
	}

	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5000" //default port
	}
	routes.RegisterTodoRoutes(app)

	log.Fatal(app.Listen(":" + PORT))
}

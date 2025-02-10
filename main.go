package main

import (
	"log"

	"github.com/Darari17/user-management/fiber/v2/config"
	"github.com/Darari17/user-management/fiber/v2/controller"
	"github.com/Darari17/user-management/fiber/v2/repository"
	"github.com/Darari17/user-management/fiber/v2/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	app := fiber.New()

	api := app.Group("/api/v2")
	userController := controller.NewUserController(userService, api)
	userController.Route()

	port := ":3000"
	log.Printf("Server is running on http://localhost%s", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

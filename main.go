package main

import (
	"golang-fiber-mongodb-hrms/config"
	"golang-fiber-mongodb-hrms/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	config.ConnectDB()

	routes.EmployeeRoutes(app)

	log.Println("Server started on port 3000")

	app.Listen(":3000")

}

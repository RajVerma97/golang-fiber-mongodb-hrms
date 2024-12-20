package routes

import (
	"golang-fiber-mongodb-hrms/controllers"
	"github.com/gofiber/fiber/v2"
)

func EmployeeRoutes(app *fiber.App) {

	api := app.Group("/api/employees")

	api.Post("/", controllers.CreateEmployee)
	api.Get("/", controllers.GetEmployees)
	api.Get("/:id", controllers.GetEmployee)
	api.Put("/:id", controllers.UpdateEmployee)
	api.Delete("/:id", controllers.DeleteEmployee)
	api.Post("/bulk", controllers.CreateMultipleEmployees)

}

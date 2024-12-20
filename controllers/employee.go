package controllers

import (
	"context"
	"golang-fiber-mongodb-hrms/config"
	"golang-fiber-mongodb-hrms/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateEmployee handles POST request to create a new employee
func CreateEmployee(c *fiber.Ctx) error {
	var employee models.Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	collection := config.Database.Collection("employees")
	employee.ID = primitive.NewObjectID()
	employee.ID = primitive.NewObjectID()
	employee.CreatedAt = time.Now()
	employee.UpdatedAt = time.Now()

	_, err := collection.InsertOne(context.Background(), employee)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(employee)
}

func CreateMultipleEmployees(c *fiber.Ctx) error {
	var employees []models.Employee
	if err := c.BodyParser(&employees); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	for i := range employees {
		employees[i].ID = primitive.NewObjectID()
		employees[i].CreatedAt = time.Now()
		employees[i].UpdatedAt = time.Now()
	}

	// Insert many employees
	collection := config.Database.Collection("employees")
	var documents []interface{}
	for _, employee := range employees {
		documents = append(documents, employee)
	}

	_, err := collection.InsertMany(context.Background(), documents)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":   "Employees created successfully",
		"employees": employees,
	})
}

// GetEmployees handles GET request to fetch all employees
func GetEmployees(c *fiber.Ctx) error {
	collection := config.Database.Collection("employees")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer cursor.Close(context.Background())

	var employees []models.Employee
	for cursor.Next(context.Background()) {
		var employee models.Employee
		if err := cursor.Decode(&employee); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		employees = append(employees, employee)
	}

	if err := cursor.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(employees)
}

// GetEmployee handles GET request to fetch a single employee by ID
func GetEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	collection := config.Database.Collection("employees")
	var employee models.Employee
	err = collection.FindOne(context.Background(), bson.M{"_id": employeeID}).Decode(&employee)
	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Employee not found"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(employee)
}

// UpdateEmployee handles PUT request to update an employee
func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var updatedEmployee models.Employee
	if err := c.BodyParser(&updatedEmployee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	collection := config.Database.Collection("employees")
	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{"_id": employeeID},
		bson.M{"$set": updatedEmployee},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Employee updated successfully"})
}

// DeleteEmployee handles DELETE request to delete an employee by ID
func DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	collection := config.Database.Collection("employees")
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": employeeID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Employee deleted successfully"})
}

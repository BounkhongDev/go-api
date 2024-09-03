package main

import (
	"encoding/json"
	"fmt"
	"go-api/config"
	"go-api/database"
	"go-api/logs"
	"go-api/routes"
	"go-api/src/controllers"
	"go-api/src/repositories"
	"go-api/src/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	//connect database
	postgresConnection, err := database.PostgresConnection()
	if err != nil {
		logs.Error(err)
		return
	}

	//basic structure
	newRepository := repositories.NewExampleRepository(postgresConnection)
	usersRepository := repositories.NewUsersRepository(postgresConnection)
	rolesRepository := repositories.NewRolesRepository(postgresConnection)

	newService := services.NewExampleService(newRepository)
	usersService := services.NewUsersService(usersRepository)
	rolesService := services.NewRolesService(rolesRepository)

	// connect route
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(logger.New())
	app.Use(cors.New())

	//example routes
	newExampleController := controllers.NewExampleController(newService)

	//users routes
	newUsersController := controllers.NewUsersController(usersService)
	newRolesController := controllers.NewRolesController(rolesService)

	newRoute := routes.NewFiberRoutes(
		//new web controller
		newExampleController,
		newUsersController,
		newRolesController,
	)
	newRoute.Install(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env("app.port"))))
}

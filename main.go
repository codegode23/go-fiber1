package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/codegode23/go-fiber1/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/codegode23/go-fiber1/docs" // you need to update github.com/rizalgowandy/go-swag-sample with your own project path
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server like that.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http

func main() {
	app := api.SetupRoute()
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Welcome to Go Fiber API")
	// })

	// app.Listen(":3000")

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New())

	// Routes
	app.Get("/", HealthCheck)
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	//app.Get("/docs/*any", swagger.HandlerDefault)

	// Start Server
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}

}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}

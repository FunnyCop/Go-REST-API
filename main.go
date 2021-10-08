package main

import (
	"github.com/gofiber/fiber/v2"
	"lbprunty.xyz/api/v2/config"
	"lbprunty.xyz/api/v2/routes"
)

func main() {

	// MySQL Connection
	db := config.MysqlConnect()

	// Fiber App
	app := fiber.New()

	routes.JobApplications(app, db)

	// Close mySQL Connection when main() is stopped
	defer db.Close()

	// Start Server on Port 3000
	app.Listen(":3000")

}

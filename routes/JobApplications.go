package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"lbprunty.xyz/api/v2/controllers"
)

func JobApplications(app *fiber.App, db *sql.DB) {

	app.Get("/api/jobApplications", func(c *fiber.Ctx) error {

		return controllers.GetJobApplications(db, c)

	})

}

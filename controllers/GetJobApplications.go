package controllers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"lbprunty.xyz/api/v2/models"
)

// Get All Job Applications
func GetJobApplications(db *sql.DB, c *fiber.Ctx) error {

	var jobApplications []models.JobApplication
	results, err := db.Query("SELECT * FROM job_applications")

	if err != nil {

		panic(err.Error())

	}

	for results.Next() {

		var jobApplication models.JobApplication

		err = results.Scan(
			&jobApplication.ID,
			&jobApplication.CompanyName,
			&jobApplication.CompanyWebsite,
			&jobApplication.CreatedAt,
			&jobApplication.Description,
			&jobApplication.Interviews,
			&jobApplication.JobTitle,
			&jobApplication.Link,
			&jobApplication.Offer,
			&jobApplication.Status,
			&jobApplication.UpdatedAt,
		)

		if err != nil {

			panic(err.Error())

		}

		jobApplications = append(jobApplications, jobApplication)

	}

	return c.JSON(jobApplications)

}

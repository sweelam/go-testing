package database

import (
	"errors"
	"log"

	model "database.learning/start/internal/models"
)

func FindAllJobs() ([]model.MyJob, error) {
	var jobs []model.MyJob
	result := db.Find(&jobs)

	if result.Error != nil {
		log.Fatal(errors.New("error while fetching my_job data"))
	}

	return jobs, result.Error
}

func GetJobById(id int) model.MyJob {
	var job = model.MyJob{ID: 1}
	db.First(&job)
	return job
}

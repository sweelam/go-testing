package service

import (
	"fmt"

	database "database.learning/start/internal/database/native"
)

type Job struct {
	ID             int     `json:"id"`
	Title          string  `json:"title"`
	StartDate      *string `json:"startDate"`
	EndDate        *string `json:"endDate"`
	Company        string  `json:"company"`
	JobDescription string  `json:"jobDescription"`
	info_ref       string
}

func getProfiles() ([]Job, error) {
	var jobs []Job
	var j Job
	rs := database.Read("SELECT * FROM my_jobs")
	defer rs.Close()
	for rs.Next() {
		if err := rs.Scan(&j.ID,
			&j.Title,
			&j.StartDate,
			&j.EndDate,
			&j.Company,
			&j.JobDescription,
			&j.info_ref); err != nil {
			return []Job{}, fmt.Errorf("error: %v", err)
		}
		jobs = append(jobs, j)
	}
	return jobs, nil
}

func GetJobs() ([]Job, error) {
	return getProfiles()
}

func DisplayJobs() {
	println("Displaying jobs")

	jobs, err := getProfiles()
	if err != nil {
		fmt.Println(err)
	}

	for _, job := range jobs {
		fmt.Println(job)
	}
}

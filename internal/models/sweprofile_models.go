package model

import "database/sql"

type MyJob struct {
	ID             int
	Title          string
	StartDate      sql.NullString
	EndDate        sql.NullString
	Company        string
	JobDescription string
	InfoRef        int
}

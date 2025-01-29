package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	AllStudents() ([]*models.Student, error)
	Connection() *sql.DB
}

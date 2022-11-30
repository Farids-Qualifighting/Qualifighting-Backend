package services

import "qualifighting.backend.de/models"

type StudentService interface {
	CreateStudent(*models.Student) error
	GetStudent(*string) (*models.Student, error)
	GetAll() ([]*models.Student, error)
	UpdateStudent(*models.Student) error
	DeleteStudent(*string) error
}

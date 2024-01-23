package services

import "github.com/Twsouza/job-rule-engine/domain"

type OptiiApiInterface interface {
	GetDepartmentByID(id int64) (*domain.Department, error)
	GetJobItemByID(id int64) (*domain.JobItem, error)
	GetLocationsByIds(id []int64) ([]domain.Location, error)
}

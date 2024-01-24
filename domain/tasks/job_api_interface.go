package tasks

import "github.com/Twsouza/job-rule-engine/domain"

type JobAPI interface {
	CreateJob(job *domain.Job) (interface{}, error)
	GetFloorRooms(floorID int) ([]domain.Location, error)
	GetFloorLocations(floorID int) ([]domain.Location, error)
}

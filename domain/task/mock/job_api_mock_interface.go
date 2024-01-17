package mock

import "github.com/Twsouza/job-rule-engine/domain"

type JobAPIMock struct {
	CreateJobFunc         func(job domain.Job) (interface{}, error)
	GetFloorRoomsFunc     func(floorID int) ([]domain.Location, error)
	GetFloorLocationsFunc func(floorID int) ([]domain.Location, error)
}

func (m *JobAPIMock) CreateJob(job domain.Job) (interface{}, error) {
	return m.CreateJobFunc(job)
}

func (m *JobAPIMock) GetFloorRooms(floorID int) ([]domain.Location, error) {
	return m.GetFloorRoomsFunc(floorID)
}

func (m *JobAPIMock) GetFloorLocations(floorID int) ([]domain.Location, error) {
	return m.GetFloorLocationsFunc(floorID)
}

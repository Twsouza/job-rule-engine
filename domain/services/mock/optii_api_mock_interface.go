package mock

import (
	"github.com/Twsouza/job-rule-engine/domain"
)

type OptiiApiMock struct {
	GetDepartmentByIDFunc func(id int64) (*domain.Department, error)
	GetJobItemByIDFunc    func(id int64) (*domain.JobItem, error)
	GetLocationsByIdsFunc func(id []int64) ([]domain.Location, error)
}

func (m *OptiiApiMock) GetDepartmentByID(id int64) (*domain.Department, error) {
	return m.GetDepartmentByIDFunc(id)
}

func (m *OptiiApiMock) GetJobItemByID(id int64) (*domain.JobItem, error) {
	return m.GetJobItemByIDFunc(id)
}

func (m *OptiiApiMock) GetLocationsByIds(id []int64) ([]domain.Location, error) {
	return m.GetLocationsByIdsFunc(id)
}

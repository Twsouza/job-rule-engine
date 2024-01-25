package roomservice

import (
	"errors"
	"testing"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/domain/tasks/mock"
	"github.com/stretchr/testify/assert"
)

func TestDeliverJobItemRoom_AssertRule(t *testing.T) {
	dj := &DeliverJobItemRoomTask{}

	t.Run("should return true for valid job request", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Room Service",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Food",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		result := dj.AssertRule(jobRequest)
		assert.True(t, result)
	})

	t.Run("should return false for invalid department", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Food",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		result := dj.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for missing job item", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Room Service",
			},
			JobItem: nil,
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		result := dj.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for invalid location", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Room Service",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Food",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Kitchen",
					},
				},
			},
		}

		result := dj.AssertRule(jobRequest)
		assert.False(t, result)
	})
}

func TestDeliverJobItemRoom_Execute(t *testing.T) {
	dj := &DeliverJobItemRoomTask{}

	t.Run("should return job result with request and result when API call is successful", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				ID:   1,
				Name: "Room Service",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Food",
			},
			Locations: []domain.Location{
				{
					ID: 1,
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		expectedJobResult := domain.JobResult{
			Request: &jobRequest,
			Result:  "success",
		}

		mockAPI := &mock.JobAPIMock{}
		mockAPI.CreateJobFunc = func(job *domain.Job) (interface{}, error) {
			return "success", nil
		}
		mockAPI.GetFloorRoomsFunc = func(locationID int) ([]domain.Location, error) {
			return []domain.Location{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			}, nil
		}

		dj.API = mockAPI

		actualJobResult := dj.Execute(jobRequest)
		assert.Equal(t, expectedJobResult, actualJobResult)
	})

	t.Run("should return job result with error when API call fails", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				ID:   1,
				Name: "Room Service",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Food",
			},
			Locations: []domain.Location{
				{
					ID: 1,
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		expectedJobResult := domain.JobResult{
			Request: &jobRequest,
			Err:     "API error",
		}

		mockAPI := &mock.JobAPIMock{}
		mockAPI.CreateJobFunc = func(job *domain.Job) (interface{}, error) {
			return nil, nil
		}
		mockAPI.GetFloorRoomsFunc = func(locationID int) ([]domain.Location, error) {
			return nil, errors.New("API error")
		}

		dj.API = mockAPI

		actualJobResult := dj.Execute(jobRequest)
		assert.Equal(t, expectedJobResult, actualJobResult)
	})
}

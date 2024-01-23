package housekeeping

import (
	"errors"
	"testing"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/domain/task/mock"
	"github.com/stretchr/testify/assert"
)

func TestCleanBedsFloor_AssertRule(t *testing.T) {
	cr := &CleanBedsFloor{}

	t.Run("should return true for valid job request", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Housekeeping",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Sheets",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		result := cr.AssertRule(jobRequest)
		assert.True(t, result)
	})

	t.Run("should return false for invalid department", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Sheets",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		result := cr.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for invalid job item", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Housekeeping",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Pillow",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		result := cr.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for empty locations", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Housekeeping",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Sheets",
			},
			Locations: []domain.Location{},
		}

		result := cr.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for missing floor location", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Housekeeping",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Sheets",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Room",
					},
				},
			},
		}

		result := cr.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for missing job item", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Housekeeping",
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

		result := cr.AssertRule(jobRequest)
		assert.False(t, result)
	})
}
func TestCleanBedsFloor_Execute(t *testing.T) {
	cr := &CleanBedsFloor{}

	t.Run("should execute clean beds floor task successfully", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				ID:   1,
				Name: "Housekeeping",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Sheets",
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

		expectedJob := domain.Job{
			Action: "clean",
			Department: domain.JDepartment{
				ID: 1,
			},
			Item: domain.JItem{
				Name: "Sheets",
			},
			Locations: []domain.JLocation{
				{
					ID: 1,
				},
			},
		}

		expectedResult := domain.JobResult{
			Request: &jobRequest,
			Result:  "job created",
		}

		mockAPI := &mock.JobAPIMock{}
		mockAPI.CreateJobFunc = func(job domain.Job) (interface{}, error) {
			assert.Equal(t, expectedJob, job)
			return "job created", nil
		}
		mockAPI.GetFloorRoomsFunc = func(floorID int) ([]domain.Location, error) {
			return []domain.Location{{ID: 1}}, nil
		}

		cr.API = mockAPI

		result := cr.Execute(jobRequest)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("should return error if GetFloorRooms fails", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				ID:   1,
				Name: "Housekeeping",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Sheets",
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

		expectedError := errors.New("failed to get floor rooms")

		mockAPI := &mock.JobAPIMock{}
		mockAPI.GetFloorRoomsFunc = func(floorID int) ([]domain.Location, error) {
			return nil, expectedError
		}

		cr.API = mockAPI

		result := cr.Execute(jobRequest)
		assert.Equal(t, expectedError, result.Err)
	})

	t.Run("should return error if CreateJob fails", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				ID:   1,
				Name: "Housekeeping",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Sheets",
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

		expectedJob := domain.Job{
			Action: "clean",
			Department: domain.JDepartment{
				ID: 1,
			},
			Item: domain.JItem{
				Name: "Sheets",
			},
			Locations: []domain.JLocation{
				{
					ID: 1,
				},
			},
		}

		expectedError := errors.New("failed to create job")

		mockAPI := &mock.JobAPIMock{}
		mockAPI.GetFloorRoomsFunc = func(floorID int) ([]domain.Location, error) {
			return []domain.Location{{ID: 1}}, nil
		}
		mockAPI.CreateJobFunc = func(job domain.Job) (interface{}, error) {
			assert.Equal(t, expectedJob, job)
			return "", expectedError
		}

		cr.API = mockAPI

		result := cr.Execute(jobRequest)
		assert.Equal(t, expectedError, result.Err)
	})
}

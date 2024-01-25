package engineering

import (
	"errors"
	"testing"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/domain/tasks/mock"
	"github.com/stretchr/testify/assert"
)

func TestRepairJobItemFloor_AssertRule(t *testing.T) {
	rj := &RepairJobItemFloor{}

	t.Run("should return true for valid job request", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Air Conditioning",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		result := rj.AssertRule(jobRequest)
		assert.True(t, result)
	})

	t.Run("should return false for invalid department", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Housekeeping",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Air Conditioning",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
			},
		}

		result := rj.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for empty job item", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
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

		result := rj.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for invalid location", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Air Conditioning",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Room",
					},
				},
			},
		}

		result := rj.AssertRule(jobRequest)
		assert.False(t, result)
	})

	t.Run("should return false for multiple locations", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Air Conditioning",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Floor",
					},
				},
				{
					LocationType: &domain.LocationType{
						DisplayName: "Room",
					},
				},
			},
		}

		result := rj.AssertRule(jobRequest)
		assert.False(t, result)
	})
}

func TestRepairJobItemFloor_Execute(t *testing.T) {
	rj := &RepairJobItemFloor{}

	t.Run("should execute repair job for valid job request", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				ID:   123,
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Air Conditioning",
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

		expectedJob := &domain.Job{
			Action: "repair",
			Department: domain.JDepartment{
				ID: 123,
			},
			Item: domain.JItem{
				Name: "Air Conditioning",
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
			Err:     "",
		}

		mockAPI := &mock.JobAPIMock{}
		mockAPI.CreateJobFunc = func(job *domain.Job) (interface{}, error) {
			assert.Equal(t, expectedJob, job)
			return "job created", nil
		}
		mockAPI.GetFloorLocationsFunc = func(floorID int) ([]domain.Location, error) {
			return []domain.Location{{ID: 1}}, nil
		}

		rj.API = mockAPI

		result := rj.Execute(jobRequest)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("should return error if GetFloorLocations fails", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				ID:   123,
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Air Conditioning",
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

		expectedError := errors.New("failed to get floor locations")

		mockAPI := &mock.JobAPIMock{}
		mockAPI.GetFloorLocationsFunc = func(floorID int) ([]domain.Location, error) {
			return nil, expectedError
		}

		rj.API = mockAPI

		result := rj.Execute(jobRequest)

		assert.Equal(t, expectedError.Error(), result.Err)
	})

	t.Run("should return error if CreateJob fails", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				ID:   123,
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Air Conditioning",
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

		expectedJob := &domain.Job{
			Action: "repair",
			Department: domain.JDepartment{
				ID: 123,
			},
			Item: domain.JItem{
				Name: "Air Conditioning",
			},
			Locations: []domain.JLocation{
				{
					ID: 1,
				},
			},
		}

		expectedError := errors.New("failed to create job")

		mockAPI := &mock.JobAPIMock{}
		mockAPI.CreateJobFunc = func(job *domain.Job) (interface{}, error) {
			assert.Equal(t, expectedJob, job)
			return nil, expectedError
		}
		mockAPI.GetFloorLocationsFunc = func(floorID int) ([]domain.Location, error) {
			return []domain.Location{{ID: 1}}, nil
		}

		rj.API = mockAPI

		result := rj.Execute(jobRequest)

		assert.Equal(t, expectedError.Error(), result.Err)
	})
}

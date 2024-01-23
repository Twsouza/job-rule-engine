package engineering

import (
	"errors"
	"testing"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/domain/tasks/mock"
	"github.com/stretchr/testify/assert"
)

func TestRepairJobItemLocation_AssertRule(t *testing.T) {
	rj := &RepairJobItemLocation{}

	t.Run("should return true for valid job request", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "TV",
			},
			Locations: []domain.Location{
				{
					LocationType: &domain.LocationType{
						DisplayName: "Room",
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
		assert.True(t, result)
	})

	t.Run("should return false for invalid department", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Housekeeping",
			},
			JobItem: &domain.JobItem{
				DisplayName: "TV",
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

	t.Run("should return false for missing job item", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
			},
			JobItem: nil,
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

	t.Run("should return false for empty locations", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "TV",
			},
			Locations: []domain.Location{},
		}

		result := rj.AssertRule(jobRequest)
		assert.False(t, result)
	})
}

func TestRepairJobItemLocation_Execute(t *testing.T) {
	rj := &RepairJobItemLocation{}

	t.Run("should execute repair job with valid job request", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				ID:   1,
				Name: "Engineering",
			},
			JobItem: &domain.JobItem{
				DisplayName: "TV",
			},
			Locations: []domain.Location{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},
		}

		expectedJob := domain.Job{
			Action: "repair",
			Department: domain.JDepartment{
				ID: 1,
			},
			Item: domain.JItem{
				Name: "TV",
			},
			Locations: []domain.JLocation{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},
		}

		expectedResult := domain.JobResult{
			Request: &jobRequest,
			Result:  "success",
			Err:     nil,
		}

		mockAPI := &mock.JobAPIMock{}
		mockAPI.CreateJobFunc = func(job domain.Job) (interface{}, error) {
			assert.Equal(t, expectedJob, job)
			return "success", nil
		}

		rj.API = mockAPI

		result := rj.Execute(jobRequest)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("should return error for invalid job request", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				ID:   2,
				Name: "Housekeeping",
			},
			JobItem: &domain.JobItem{
				DisplayName: "TV",
			},
			Locations: []domain.Location{
				{
					ID: 1,
				},
			},
		}

		expectedJob := domain.Job{
			Action: "repair",
			Department: domain.JDepartment{
				ID: 2,
			},
			Item: domain.JItem{
				Name: "TV",
			},
			Locations: []domain.JLocation{
				{
					ID: 1,
				},
			},
		}

		expectedResult := domain.JobResult{
			Request: &jobRequest,
			Result:  "",
			Err:     errors.New("failed to create job"),
		}

		mockAPI := &mock.JobAPIMock{}
		mockAPI.CreateJobFunc = func(job domain.Job) (interface{}, error) {
			assert.Equal(t, expectedJob, job)
			return "", errors.New("failed to create job")
		}

		rj.API = mockAPI

		result := rj.Execute(jobRequest)
		assert.Equal(t, expectedResult, result)
	})
}

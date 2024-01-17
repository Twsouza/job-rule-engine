package task

import (
	"errors"
	"testing"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/domain/task/mock"
	"github.com/stretchr/testify/assert"
)

func TestDeliverJobItemLocation_AssertRule(t *testing.T) {
	dj := &DeliverJobItemLocationTask{}

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

	t.Run("should return false for empty locations", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				Name: "Room Service",
			},
			JobItem: &domain.JobItem{
				DisplayName: "Food",
			},
			Locations: []domain.Location{},
		}

		result := dj.AssertRule(jobRequest)
		assert.False(t, result)
	})
}

func TestDeliverJobItemLocation_Execute(t *testing.T) {
	dj := &DeliverJobItemLocationTask{}

	t.Run("should execute job request and return job result", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				ID: 1,
			},
			JobItem: &domain.JobItem{
				DisplayName: "Food",
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
			Action: "deliver",
			Department: domain.JDepartment{
				ID: 1,
			},
			Item: domain.JItem{
				Name: "Food",
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
			Result:  "job result",
			Err:     nil,
		}

		mockAPI := &mock.JobAPIMock{}
		mockAPI.CreateJobFunc = func(job domain.Job) (interface{}, error) {
			assert.Equal(t, expectedJob, job)
			return "job result", nil
		}

		dj.Api = mockAPI

		result := dj.Execute(jobRequest)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("should execute job request and return error", func(t *testing.T) {
		jobRequest := domain.JobRequest{
			Department: &domain.Department{
				ID: 1,
			},
			JobItem: &domain.JobItem{
				DisplayName: "Food",
			},
			Locations: []domain.Location{
				{
					ID: 1,
				},
			},
		}

		expectedJob := domain.Job{
			Action: "deliver",
			Department: domain.JDepartment{
				ID: 1,
			},
			Item: domain.JItem{
				Name: "Food",
			},
			Locations: []domain.JLocation{
				{
					ID: 1,
				},
			},
		}

		expectedError := errors.New("job creation failed")

		mockAPI := &mock.JobAPIMock{}
		mockAPI.CreateJobFunc = func(job domain.Job) (interface{}, error) {
			assert.Equal(t, expectedJob, job)
			return nil, expectedError
		}

		dj.Api = mockAPI

		result := dj.Execute(jobRequest)
		assert.Equal(t, expectedError, result.Err)
	})
}

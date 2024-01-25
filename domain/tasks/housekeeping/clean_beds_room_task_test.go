package housekeeping

import (
	"testing"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/domain/tasks/mock"
	"github.com/stretchr/testify/assert"
)

func TestCleanBedsRoom_AssertRule(t *testing.T) {
	cr := &CleanBedsRoom{}

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
						DisplayName: "Room",
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
						DisplayName: "Room",
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
						DisplayName: "Room",
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

	t.Run("should return false for missing room location type", func(t *testing.T) {
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
						DisplayName: "Bathroom",
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
						DisplayName: "Room",
					},
				},
			},
		}

		result := cr.AssertRule(jobRequest)
		assert.False(t, result)
	})
}

func TestCleanBedsRoom_Execute(t *testing.T) {
	cr := &CleanBedsRoom{}

	t.Run("should execute clean beds room task", func(t *testing.T) {
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
						DisplayName: "Room",
					},
				},
			},
		}

		expectedJob := &domain.Job{
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
			Result:  "success",
			Err:     "",
		}

		mockAPI := &mock.JobAPIMock{}
		mockAPI.CreateJobFunc = func(job *domain.Job) (interface{}, error) {
			assert.Equal(t, expectedJob, job)
			return "success", nil
		}

		cr.API = mockAPI

		result := cr.Execute(jobRequest)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("should not execute clean beds room task for invalid location type", func(t *testing.T) {
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
						DisplayName: "Bathroom",
					},
				},
			},
		}

		expectedResult := domain.JobResult{
			Request: &jobRequest,
			Result:  nil,
			Err:     "no locations found for this job",
		}

		mockAPI := &mock.JobAPIMock{}
		mockAPI.CreateJobFunc = func(job *domain.Job) (interface{}, error) {
			t.Error("CreateJob should not be called")
			return "", nil
		}

		cr.API = mockAPI

		result := cr.Execute(jobRequest)
		assert.Equal(t, expectedResult, result)
	})
}

type MockAPI struct {
	CreateJobFunc func(job domain.Job) (string, error)
}

func (m *MockAPI) CreateJob(job domain.Job) (string, error) {
	return m.CreateJobFunc(job)
}

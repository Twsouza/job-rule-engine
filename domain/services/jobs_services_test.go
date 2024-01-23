package services

import (
	"errors"
	"testing"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/domain/tasks"
	"github.com/stretchr/testify/assert"
)

type MockRule struct {
	AssertFunc  func(jobRequest domain.JobRequest) bool
	ExecuteFunc func(jobRequest domain.JobRequest) domain.JobResult
}

func (mr *MockRule) AssertRule(jobRequest domain.JobRequest) bool {
	return mr.AssertFunc(jobRequest)
}

func (mr *MockRule) Execute(jobRequest domain.JobRequest) domain.JobResult {
	return mr.ExecuteFunc(jobRequest)
}

func TestCreateJob(t *testing.T) {
	// Creates 2 rules that will execute concurrently
	mockRules := []tasks.JobTask{
		&MockRule{
			AssertFunc: func(jobRequest domain.JobRequest) bool {
				return jobRequest.Department.Name == "Engineering"
			},
			ExecuteFunc: func(jobRequest domain.JobRequest) domain.JobResult {
				return domain.JobResult{}
			},
		},
		&MockRule{
			AssertFunc: func(jobRequest domain.JobRequest) bool {
				return jobRequest.Department.Name == "Engineering"
			},
			ExecuteFunc: func(jobRequest domain.JobRequest) domain.JobResult {
				return domain.JobResult{}
			},
		},
		&MockRule{
			AssertFunc: func(jobRequest domain.JobRequest) bool {
				return jobRequest.Department.Name == "Housekeeping"
			},
			ExecuteFunc: func(jobRequest domain.JobRequest) domain.JobResult {
				return domain.JobResult{}
			},
		},
	}

	// Create an instance of JobService with the mock rules
	jobService := &JobService{
		Tasks: mockRules,
	}

	// Define a test job request
	jobRequest := &domain.JobRequest{
		Department: &domain.Department{
			Name: "Engineering",
		},
	}

	t.Run("should return no errors when all rules execute successfully", func(t *testing.T) {
		// Call the CreateJob function
		jr := jobService.CreateJob(jobRequest)
		assert.Len(t, jr, 2)
		assert.Nil(t, jr[0].Err)
		assert.Nil(t, jr[1].Err)
	})

	t.Run("should return an error when a rule fails", func(t *testing.T) {
		jobService.Tasks = append(jobService.Tasks, &MockRule{
			AssertFunc: func(jobRequest domain.JobRequest) bool {
				return jobRequest.Department.Name == "Engineering"
			},
			ExecuteFunc: func(jobRequest domain.JobRequest) domain.JobResult {
				return domain.JobResult{
					Err: errors.New("failed to execute rule"),
				}
			},
		})

		// Call the CreateJob function
		jr := jobService.CreateJob(jobRequest)
		assert.Len(t, jr, 3)
		assert.Nil(t, jr[0].Err)
		assert.Nil(t, jr[1].Err)
		assert.NotNil(t, jr[2].Err)
	})
}

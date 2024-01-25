package services

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Twsouza/job-rule-engine/application/dto"
	"github.com/Twsouza/job-rule-engine/domain"
	servicesMock "github.com/Twsouza/job-rule-engine/domain/services/mock"
	"github.com/Twsouza/job-rule-engine/domain/tasks"
	"github.com/Twsouza/job-rule-engine/domain/tasks/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreateJob(t *testing.T) {
	// Creates 2 rules that will execute concurrently
	mockRules := []tasks.JobTask{
		&mock.MockRule{
			AssertFunc: func(jobRequest domain.JobRequest) bool {
				return jobRequest.Department.Name == "Engineering"
			},
			ExecuteFunc: func(jobRequest domain.JobRequest) domain.JobResult {
				return domain.JobResult{}
			},
		},
		&mock.MockRule{
			AssertFunc: func(jobRequest domain.JobRequest) bool {
				return jobRequest.Department.Name == "Engineering"
			},
			ExecuteFunc: func(jobRequest domain.JobRequest) domain.JobResult {
				return domain.JobResult{}
			},
		},
		&mock.MockRule{
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
		assert.Empty(t, jr[0].Err)
		assert.Empty(t, jr[1].Err)
	})

	t.Run("should return an error when a rule fails", func(t *testing.T) {
		jobService.Tasks = append(jobService.Tasks, &mock.MockRule{
			AssertFunc: func(jobRequest domain.JobRequest) bool {
				return jobRequest.Department.Name == "Engineering"
			},
			ExecuteFunc: func(jobRequest domain.JobRequest) domain.JobResult {
				return domain.JobResult{
					Err: "failed to execute rule",
				}
			},
		})

		// Call the CreateJob function
		jr := jobService.CreateJob(jobRequest)
		assert.Len(t, jr, 3)
		assert.Contains(t, jr, domain.JobResult{Err: "failed to execute rule"})
	})
}

func TestLoadJob(t *testing.T) {
	optiiAPIMock := &servicesMock.OptiiApiMock{}

	jobService := &JobService{
		OptiiAPI: optiiAPIMock,
	}

	reqDto := &dto.JobRequestDto{
		DepartmentID: 1,
		JobItemID:    2,
		LocationsID:  []int64{3, 4, 5},
	}

	department := &domain.Department{
		ID:   1,
		Name: "Engineering",
	}
	jobItem := &domain.JobItem{
		ID:          1,
		DisplayName: "Job Item",
	}
	locations := []domain.Location{
		{
			ID:          1,
			Name:        "1",
			DisplayName: "Location 1",
			ParentLocation: &domain.ParentLocation{
				ID:          2,
				Name:        "2",
				DisplayName: "Location 2",
			},
			LocationType: &domain.LocationType{
				ID:          3,
				DisplayName: "Room",
			},
		},
	}

	t.Run("should load job request successfully", func(t *testing.T) {
		optiiAPIMock.GetDepartmentByIDFunc = func(id int64) (*domain.Department, error) {
			assert.Equal(t, reqDto.DepartmentID, id)
			return department, nil
		}

		optiiAPIMock.GetJobItemByIDFunc = func(id int64) (*domain.JobItem, error) {
			assert.Equal(t, reqDto.JobItemID, id)
			return jobItem, nil
		}

		optiiAPIMock.GetLocationsByIdsFunc = func(ids []int64) ([]domain.Location, error) {
			assert.Equal(t, reqDto.LocationsID, ids)
			return locations, nil
		}
		expectedJobRequest := &domain.JobRequest{
			Department: department,
			JobItem:    jobItem,
			Locations:  locations,
		}

		jr, err := jobService.LoadJob(reqDto)
		assert.Len(t, err, 0)
		assert.Equal(t, expectedJobRequest, jr)
	})

	t.Run("should return an error if GetDepartmentByID fails", func(t *testing.T) {
		departmentError := errors.New("failed to get department")
		expectedError := []error{fmt.Errorf("department %w", departmentError)}

		optiiAPIMock.GetDepartmentByIDFunc = func(id int64) (*domain.Department, error) {
			assert.Equal(t, reqDto.DepartmentID, id)
			return nil, departmentError
		}

		optiiAPIMock.GetJobItemByIDFunc = func(id int64) (*domain.JobItem, error) {
			assert.Equal(t, reqDto.JobItemID, id)
			return jobItem, nil
		}

		optiiAPIMock.GetLocationsByIdsFunc = func(ids []int64) ([]domain.Location, error) {
			assert.Equal(t, reqDto.LocationsID, ids)
			return locations, nil
		}

		jr, err := jobService.LoadJob(reqDto)
		assert.Len(t, err, 1)
		assert.Error(t, err[0])
		assert.Nil(t, jr.Department)
		assert.Equal(t, expectedError, err)
	})

	t.Run("should return an error if GetJobItemByID fails", func(t *testing.T) {
		jobItemError := errors.New("failed to get job item")
		expectedError := []error{fmt.Errorf("jobItem %w", jobItemError)}

		optiiAPIMock.GetDepartmentByIDFunc = func(id int64) (*domain.Department, error) {
			assert.Equal(t, reqDto.DepartmentID, id)
			return department, nil
		}

		optiiAPIMock.GetJobItemByIDFunc = func(id int64) (*domain.JobItem, error) {
			assert.Equal(t, reqDto.JobItemID, id)
			return nil, jobItemError
		}

		optiiAPIMock.GetLocationsByIdsFunc = func(ids []int64) ([]domain.Location, error) {
			assert.Equal(t, reqDto.LocationsID, ids)
			return locations, nil
		}

		jr, err := jobService.LoadJob(reqDto)
		assert.Len(t, err, 1)
		assert.Error(t, err[0])
		assert.Nil(t, jr.JobItem)
		assert.Equal(t, expectedError, err)
	})

	t.Run("should return an error if GetLocationsByIds fails", func(t *testing.T) {
		getLocationsError := errors.New("failed to get locations")
		expectedError := []error{fmt.Errorf("location %w", getLocationsError)}

		optiiAPIMock.GetDepartmentByIDFunc = func(id int64) (*domain.Department, error) {
			assert.Equal(t, reqDto.DepartmentID, id)
			return department, nil
		}

		optiiAPIMock.GetJobItemByIDFunc = func(id int64) (*domain.JobItem, error) {
			assert.Equal(t, reqDto.JobItemID, id)
			return jobItem, nil
		}

		optiiAPIMock.GetLocationsByIdsFunc = func(ids []int64) ([]domain.Location, error) {
			assert.Equal(t, reqDto.LocationsID, ids)
			return nil, getLocationsError
		}

		jr, err := jobService.LoadJob(reqDto)
		assert.Len(t, err, 1)
		assert.Error(t, err[0])
		assert.Len(t, jr.Locations, 0)
		assert.Equal(t, expectedError, err)
	})
}

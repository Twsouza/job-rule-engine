package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Twsouza/job-rule-engine/application/dto"
	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/domain/services/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateJob(t *testing.T) {
	t.Run("should return status created when job is created", func(t *testing.T) {
		// Create a new Gin router
		router := gin.Default()

		mockJobService := &mock.JobServiceMock{}
		mockJobService.CreateJobFunc = func(jobRequest *domain.JobRequest) []domain.JobResult {
			return []domain.JobResult{
				{
					Request: jobRequest,
					Result:  "success",
					Err:     "",
				},
			}
		}
		mockJobService.LoadJobFunc = func(dto *dto.JobRequestDto) (*domain.JobRequest, []error) {
			jr := &domain.JobRequest{
				Department: &domain.Department{
					ID:   1,
					Name: "Engineering",
				},
				JobItem: &domain.JobItem{
					ID:          1,
					DisplayName: "Item",
				},
				Locations: []domain.Location{
					{
						ID:          1,
						Name:        "Location",
						DisplayName: "Location 1",
					},
				},
			}
			return jr, nil
		}

		// Create a new instance of JobRuleEngineHandler
		handler := &JobRuleEngineHandler{
			JobService: mockJobService,
		}

		// Define a test request body
		reqBody := `{"departmentId": 1, "jobItemId": 1, "locationsId": [1]}`

		// Create a new HTTP request with the test request body
		req, err := http.NewRequest("POST", "/jobs", strings.NewReader(reqBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		// Create a new HTTP response recorder
		res := httptest.NewRecorder()

		// Set up the Gin router to handle the request
		router.POST("/jobs", handler.CreateJob)

		// Perform the request
		router.ServeHTTP(res, req)

		// Assert the response status code
		assert.Equal(t, http.StatusOK, res.Code)

		// Assert the response body
		expectedBody := `[{"request":{"department":{"id":1,"name":"Engineering"},"jobItem":{"id":1,"displayName":"Item"},"locations":[{"id":1,"name":"Location","displayName":"Location 1"}]},"result":"success","error":""}]`
		assert.Equal(t, expectedBody, res.Body.String())
	})

	t.Run("should return status bad request when job is not created", func(t *testing.T) {
		// Create a new Gin router
		router := gin.Default()

		mockJobService := &mock.JobServiceMock{}
		mockJobService.CreateJobFunc = func(jobRequest *domain.JobRequest) []domain.JobResult {
			return []domain.JobResult{}
		}
		mockJobService.LoadJobFunc = func(dto *dto.JobRequestDto) (*domain.JobRequest, []error) {
			jr := &domain.JobRequest{
				Department: &domain.Department{
					ID:   1,
					Name: "Engineering",
				},
				JobItem: &domain.JobItem{
					ID:          1,
					DisplayName: "Item",
				},
				Locations: []domain.Location{
					{
						ID:          1,
						Name:        "Location",
						DisplayName: "Location 1",
					},
				},
			}
			return jr, nil
		}

		// Create a new instance of JobRuleEngineHandler
		handler := &JobRuleEngineHandler{
			JobService: mockJobService,
		}

		// Define a test request body
		reqBody := `{"departmentId": 1, "jobItemId": 1, "locationsId": [1]}`

		// Create a new HTTP request with the test request body
		req, err := http.NewRequest("POST", "/jobs", strings.NewReader(reqBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		// Create a new HTTP response recorder
		res := httptest.NewRecorder()

		// Set up the Gin router to handle the request
		router.POST("/jobs", handler.CreateJob)

		// Perform the request
		router.ServeHTTP(res, req)

		// Assert the response status code
		assert.Equal(t, http.StatusBadRequest, res.Code)

		// Assert the response body
		expectedBody := `{"error":"no rules matched for this job"}`
		assert.Equal(t, expectedBody, res.Body.String())
	})
}

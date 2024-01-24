package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/stretchr/testify/assert"
)

func TestOptiiSdk_GetDepartmentByID(t *testing.T) {
	t.Run("should return the department when the request is successful", func(t *testing.T) {
		// Create a test department
		department := &domain.Department{
			ID:   1,
			Name: "Test Department",
		}

		// Create a mock server to handle the GET request
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Verify the request method and URL
			assert.Equal(t, http.MethodGet, r.Method)

			expectedURL := fmt.Sprintf("/api/v1/departments/%d", department.ID)
			assert.Equal(t, expectedURL, r.URL.Path)

			// Return the test department as the response
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(department)
		}))
		defer server.Close()

		// Create an instance of OptiiSdk with the mock server URL
		optiiSdk := &OptiiSdk{
			BaseUrl:    server.URL,
			ApiVersion: "v1",
			Client:     http.DefaultClient,
		}

		// Call the GetDepartmentByID method
		result, err := optiiSdk.GetDepartmentByID(1)
		assert.NoError(t, err)

		// Verify the result
		assert.Equal(t, department, result)
	})

	t.Run("should return an error when the request fails with a 404", func(t *testing.T) {
		// Create a mock server to handle the GET request
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body := `{
				"type": "https://datatracker.ietf.org/doc/html/rfc7231#section-6.5.4",
				"title": "Not Found",
				"status": 404,
				"detail": "99999 not found",
			}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(body))
		}))

		defer server.Close()

		// Create an instance of OptiiSdk with the mock server URL
		optiiSdk := &OptiiSdk{
			BaseUrl:    server.URL,
			ApiVersion: "v1",
			Client:     http.DefaultClient,
		}

		// Call the GetDepartmentByID method
		_, err := optiiSdk.GetDepartmentByID(99999)
		assert.EqualError(t, err, "Not Found: 99999 not found")
	})
}

func TestOptiiSdk_GetJobItemByID(t *testing.T) {
	t.Run("should return the job item when the request is successful", func(t *testing.T) {
		jobItem := &domain.JobItem{
			ID:          1,
			DisplayName: "Test Job Item",
		}

		// Create a mock server to handle the GET request
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Verify the request method and URL
			assert.Equal(t, http.MethodGet, r.Method)

			expectedURL := fmt.Sprintf("/api/v1/jobitems/%d", jobItem.ID)
			assert.Equal(t, expectedURL, r.URL.Path)
			// Return the test department as the response
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(jobItem)
		}))
		defer server.Close()

		// Create an instance of OptiiSdk with the mock server URL
		optiiSdk := &OptiiSdk{
			BaseUrl:    server.URL,
			ApiVersion: "v1",
			Client:     http.DefaultClient,
		}

		// Call the GetJobItemByID method
		result, err := optiiSdk.GetJobItemByID(1)
		assert.NoError(t, err)

		// Verify the result
		assert.Equal(t, jobItem, result)
	})

	t.Run("should return an error when the request fails with a 404", func(t *testing.T) {
		// Create a mock server to handle the GET request
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body := `{
				"type": "https://datatracker.ietf.org/doc/html/rfc7231#section-6.5.4",
				"title": "Not Found",
				"status": 404,
				"detail": "99999 not found",
			}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(body))
		}))

		defer server.Close()

		// Create an instance of OptiiSdk with the mock server URL
		optiiSdk := &OptiiSdk{
			BaseUrl:    server.URL,
			ApiVersion: "v1",
			Client:     http.DefaultClient,
		}

		// Call the GetJobItemByID method
		_, err := optiiSdk.GetJobItemByID(99999)
		assert.EqualError(t, err, "Not Found: 99999 not found")
	})
}

func TestOptiiSdk_GetLocationByID(t *testing.T) {
	t.Run("should return the job item when the request is successful", func(t *testing.T) {
		location := &domain.Location{
			ID:          1,
			Name:        "1",
			DisplayName: "Location 1",
			LocationType: &domain.LocationType{
				ID:          1,
				DisplayName: "Room",
			},
		}

		// Create a mock server to handle the GET request
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Verify the request method and URL
			assert.Equal(t, http.MethodGet, r.Method)

			expectedURL := fmt.Sprintf("/api/v1/locations/%d", location.ID)
			assert.Equal(t, expectedURL, r.URL.Path)
			// Return the test department as the response
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(location)
		}))
		defer server.Close()

		// Create an instance of OptiiSdk with the mock server URL
		optiiSdk := &OptiiSdk{
			BaseUrl:    server.URL,
			ApiVersion: "v1",
			Client:     http.DefaultClient,
		}

		// Call the GetLocationByID method
		result, err := optiiSdk.GetLocationByID(1)
		assert.NoError(t, err)

		// Verify the result
		assert.Equal(t, location, result)
	})

	t.Run("should return an error when the request fails with a 404", func(t *testing.T) {
		// Create a mock server to handle the GET request
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body := `{
				"type": "https://datatracker.ietf.org/doc/html/rfc7231#section-6.5.4",
				"title": "Not Found",
				"status": 404,
				"detail": "99999 not found",
			}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(body))
		}))

		defer server.Close()

		// Create an instance of OptiiSdk with the mock server URL
		optiiSdk := &OptiiSdk{
			BaseUrl:    server.URL,
			ApiVersion: "v1",
			Client:     http.DefaultClient,
		}

		// Call the GetLocationByID method
		_, err := optiiSdk.GetLocationByID(99999)
		assert.EqualError(t, err, "Not Found: 99999 not found")
	})
}

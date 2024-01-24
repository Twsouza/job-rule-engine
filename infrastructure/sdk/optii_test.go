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

func TestOptiiSdk_CreateJob(t *testing.T) {
	t.Run("should return the job item when the request is successful", func(t *testing.T) {
		job := &domain.Job{
			Item: domain.JItem{
				Name: "Test Job Item",
			},
			Department: domain.JDepartment{
				ID: 1,
			},
			Locations: []domain.JLocation{
				{
					ID: 1,
				},
			},
			Action: "deliver",
		}

		// Create a mock server to handle the GET request
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Verify the request method and URL
			assert.Equal(t, http.MethodPost, r.Method)

			receivedJob := &domain.Job{}
			err := json.NewDecoder(r.Body).Decode(receivedJob)
			assert.NoError(t, err)

			expectedURL := "/api/v1/jobs"
			assert.Equal(t, expectedURL, r.URL.Path)

			jobCreated := `{"id":4393,"item":{"displayname":"drink"},"type":"guest","priority":"medium","action":"deliver","attachments":null,"locations":[{"id":10,"name":"Dumpster Area"}],"departments":[{"id":7,"name":"Housekeeping"}],"roles":[{"id":10,"name":"IT Administrator"}],"notes":null,"assignee":null,"dueBy":"2024-01-24T17:05:48.000Z"}`

			// Return the test department as the response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(jobCreated))
		}))
		defer server.Close()

		// Create an instance of OptiiSdk with the mock server URL
		optiiSdk := &OptiiSdk{
			BaseUrl:    server.URL,
			ApiVersion: "v1",
			Client:     http.DefaultClient,
		}

		// Call the CreateJob method
		result, err := optiiSdk.CreateJob(job)
		fmt.Printf("%+v", result)
		fmt.Printf("%+v", err)
		assert.NoError(t, err)

		// Verify the result
		assert.Equal(t, 4393, result.(*domain.JobCreated).ID)
		assert.Equal(t, "deliver", result.(*domain.JobCreated).Action)
	})

	t.Run("should return an error when the request fails with a 400", func(t *testing.T) {
		// Create a mock server to handle the POST request
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body := `{
        "type": "https://datatracker.ietf.org/doc/html/rfc7231#section-6.5.4",
        "title": "Bad Request",
        "status": 400,
        "detail": "Invalid request",
      }`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(body))
		}))

		defer server.Close()

		// Create an instance of OptiiSdk with the mock server URL
		optiiSdk := &OptiiSdk{
			BaseUrl:    server.URL,
			ApiVersion: "v1",
			Client:     http.DefaultClient,
		}

		// Call the CreateJob method
		_, err := optiiSdk.CreateJob(&domain.Job{})
		assert.EqualError(t, err, "Bad Request: Invalid request")
	})
}

func TestOptiiSdk_GetFloorRooms(t *testing.T) {
	t.Run("should return the floor rooms when the request is successful", func(t *testing.T) {
		// Create a mock server to handle the GetLocations request
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Verify the request method and URL
			assert.Equal(t, http.MethodGet, r.Method)

			expectedURL := fmt.Sprintf("/api/v1/locations?first=%d&next=%d&locationType=%s", 0, 100, "Room")
			assert.Equal(t, expectedURL, r.URL.RequestURI())

			// Create a sample response with floor rooms
			locationsQuery := &LocationsQuery{
				Locations: []domain.Location{
					{
						ID:   1,
						Name: "Room 1",
						ParentLocation: &domain.ParentLocation{
							ID:   1,
							Name: "Floor 1",
						},
						LocationType: &domain.LocationType{
							ID:          1,
							DisplayName: "Room",
						},
					},
					{
						ID:   2,
						Name: "Room 2",
						ParentLocation: &domain.ParentLocation{
							ID:   1,
							Name: "Floor 1",
						},
						LocationType: &domain.LocationType{
							ID:          1,
							DisplayName: "Room",
						},
					},
					{
						ID:   3,
						Name: "Room 3",
						ParentLocation: &domain.ParentLocation{
							ID:   2,
							Name: "Floor 2",
						},
						LocationType: &domain.LocationType{
							ID:          1,
							DisplayName: "Room",
						},
					},
				},
				PageInfo: &PageInfo{
					HasNextPage: false,
					EndCursor:   100,
				},
			}

			// Return the sample response as the response
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(locationsQuery)
		}))
		defer server.Close()

		// Create an instance of OptiiSdk with the mock server URL
		optiiSdk := &OptiiSdk{
			BaseUrl:    server.URL,
			ApiVersion: "v1",
			Client:     http.DefaultClient,
		}

		// Call the GetFloorRooms method
		result, err := optiiSdk.GetFloorRooms(1)
		assert.NoError(t, err)

		// Verify the result
		expectedLocations := []domain.Location{
			{
				ID:   1,
				Name: "Room 1",
				ParentLocation: &domain.ParentLocation{
					ID:   1,
					Name: "Floor 1",
				},
				LocationType: &domain.LocationType{
					ID:          1,
					DisplayName: "Room",
				},
			},
			{
				ID:   2,
				Name: "Room 2",
				ParentLocation: &domain.ParentLocation{
					ID:   1,
					Name: "Floor 1",
				},
				LocationType: &domain.LocationType{
					ID:          1,
					DisplayName: "Room",
				},
			},
		}
		assert.Equal(t, expectedLocations, result)
		assert.Equal(t, 2, len(result))
	})

	t.Run("should return an error when the GetLocations request fails", func(t *testing.T) {
		// Create a mock server to handle the GetLocations request
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			errMsg := `{
				"type": "https://tools.ietf.org/html/rfc7231#section-6.5.1",
				"title": "An error has occured.",
				"status": 400,
				"detail": "The input was not a valid value."
			}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errMsg))
		}))
		defer server.Close()

		// Create an instance of OptiiSdk with the mock server URL
		optiiSdk := &OptiiSdk{
			BaseUrl:    server.URL,
			ApiVersion: "v1",
			Client:     http.DefaultClient,
		}

		// Call the GetFloorRooms method
		_, err := optiiSdk.GetFloorRooms(1)
		assert.Error(t, err)
		assert.Equal(t, "An error has occured.: The input was not a valid value.", err.Error())
	})
}

func TestOptiiSdk_GetFloorLocations(t *testing.T) {
	t.Run("should return the floor locations when the request is successful", func(t *testing.T) {

		// Create a mock server to handle the GetLocations request
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Verify the request method and URL
			assert.Equal(t, http.MethodGet, r.Method)

			expectedURL := fmt.Sprintf("/api/v1/locations?first=%d&next=%d", 0, 100)
			assert.Equal(t, expectedURL, r.URL.RequestURI())

			// Create a sample response with floor locations
			locationsQuery := &LocationsQuery{
				Locations: []domain.Location{
					{
						ID:   1,
						Name: "Location 1",
						ParentLocation: &domain.ParentLocation{
							ID:   1,
							Name: "Floor 1",
						},
						LocationType: &domain.LocationType{
							ID:          1,
							DisplayName: "Room",
						},
					},
					{
						ID:   2,
						Name: "Location 2",
						ParentLocation: &domain.ParentLocation{
							ID:   1,
							Name: "Floor 1",
						},
						LocationType: &domain.LocationType{
							ID:          2,
							DisplayName: "Public Area",
						},
					},
					{
						ID:   3,
						Name: "Location 3",
						ParentLocation: &domain.ParentLocation{
							ID:   2,
							Name: "Floor 2",
						},
						LocationType: &domain.LocationType{
							ID:          3,
							DisplayName: "Floor",
						},
					},
				},
				PageInfo: &PageInfo{
					HasNextPage: false,
					EndCursor:   100,
				},
			}

			// Return the sample response as the response
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(locationsQuery)
		}))
		defer server.Close()

		// Create an instance of OptiiSdk with the mock server URL
		optiiSdk := &OptiiSdk{
			BaseUrl:    server.URL,
			ApiVersion: "v1",
			Client:     http.DefaultClient,
		}

		// Call the GetFloorLocations method
		result, err := optiiSdk.GetFloorLocations(1)
		assert.NoError(t, err)

		// Verify the result
		expectedLocations := []domain.Location{
			{
				ID:   1,
				Name: "Location 1",
				ParentLocation: &domain.ParentLocation{
					ID:   1,
					Name: "Floor 1",
				},
				LocationType: &domain.LocationType{
					ID:          1,
					DisplayName: "Room",
				},
			},
			{
				ID:   2,
				Name: "Location 2",
				ParentLocation: &domain.ParentLocation{
					ID:   1,
					Name: "Floor 1",
				},
				LocationType: &domain.LocationType{
					ID:          2,
					DisplayName: "Public Area",
				},
			},
		}
		assert.Equal(t, expectedLocations, result)
		assert.Equal(t, 2, len(result))
	})

	t.Run("should return an error when the request fails", func(t *testing.T) {
		// Create a mock server to handle the GetLocations request
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			errMsg := `{
				"type": "https://tools.ietf.org/html/rfc7231#section-6.5.1",
				"title": "An error has occured.",
				"status": 400,
				"detail": "The input was not a valid value."
			}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errMsg))
		}))
		defer server.Close()

		// Create an instance of OptiiSdk with the mock server URL
		optiiSdk := &OptiiSdk{
			BaseUrl:    server.URL,
			ApiVersion: "v1",
			Client:     http.DefaultClient,
		}

		// Call the GetFloorLocations method
		_, err := optiiSdk.GetFloorLocations(1)
		assert.Error(t, err)
		assert.Equal(t, "An error has occured.: The input was not a valid value.", err.Error())
	})
}

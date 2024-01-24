package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/infrastructure/pkg"
	"github.com/hashicorp/go-retryablehttp"
)

// HTTPClientInterface defines the interface for HTTP client operations
type HTTPClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

type OptiiSdk struct {
	BaseUrl    string
	ApiVersion string
	RetryMax   int
	Client     HTTPClientInterface
}

func NewOptiiSdk(baseUrl, apiVersion string, retryMax int, httpClientInterface *HTTPClientInterface) (*OptiiSdk, error) {
	if baseUrl == "" {
		return nil, fmt.Errorf("baseUrl is required")
	}

	if apiVersion == "" {
		apiVersion = "v1"
	}

	if retryMax == 0 {
		retryMax = 3
	}

	optii := &OptiiSdk{
		BaseUrl:    baseUrl,
		ApiVersion: apiVersion,
		RetryMax:   retryMax,
	}

	if httpClientInterface == nil {
		client, err := RetryableHttpClient(retryMax)
		if err != nil {
			return nil, err
		}
		optii.Client = client
	}

	return optii, nil
}

func RetryableHttpClient(retryMax int) (*http.Client, error) {
	// Create a new retryable HTTP client
	client := retryablehttp.NewClient()
	client.RetryMax = retryMax

	auth, err := pkg.Authenticate()
	if err != nil {
		return nil, err
	}

	// RequestLogHook allows a user-supplied function to be called before each retry.
	// It will be used to set the Authorization header
	client.RequestLogHook = func(logger retryablehttp.Logger, request *http.Request, retryNumber int) {
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", auth.AccessToken))
	}

	return client.StandardClient(), nil
}

// GetDepartmentByID retrieves a department by its ID.
func (o *OptiiSdk) GetDepartmentByID(id int64) (*domain.Department, error) {
	// Create a new GET request
	endpoint := fmt.Sprintf("%s/api/%s/departments/%d", o.BaseUrl, o.ApiVersion, id)
	request, err := retryablehttp.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Send the request
	response, err := o.Client.Do(request.Request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Parse the response body
	var department domain.Department
	err = ParseResponse(response, &department)
	if err != nil {
		return nil, err
	}

	return &department, nil
}

// GetJobItemByID retrieves a job item by its ID from the Optii SDK.
func (o *OptiiSdk) GetJobItemByID(id int64) (*domain.JobItem, error) {
	// Create a new GET request
	endpoint := fmt.Sprintf("%s/api/%s/jobitems/%d", o.BaseUrl, o.ApiVersion, id)
	request, err := retryablehttp.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Send the request
	response, err := o.Client.Do(request.Request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Parse the response body
	var jobItem domain.JobItem
	err = ParseResponse(response, &jobItem)
	if err != nil {
		return nil, err
	}

	return &jobItem, nil
}

// GetLocationsByIds retrieves locations by their IDs.
func (o *OptiiSdk) GetLocationsByIds(ids []int64) ([]domain.Location, error) {
	wg := sync.WaitGroup{}
	locationChan := make(chan domain.Location, len(ids))
	errChan := make(chan error)

	for _, id := range ids {
		wg.Add(1)
		go func(id int64) {
			defer wg.Done()
			location, err := o.GetLocationByID(id)
			if err != nil {
				errChan <- err
				return
			}
			locationChan <- *location
		}(id)
	}

	go func() {
		wg.Wait()
		close(locationChan)
		close(errChan)
	}()

	var locations []domain.Location
	for location := range locationChan {
		locations = append(locations, location)
	}

	for err := range errChan {
		return nil, err
	}

	return locations, nil
}

// GetLocationByID retrieves a location by its ID.
func (o *OptiiSdk) GetLocationByID(id int64) (*domain.Location, error) {
	// Create a new GET request
	endpoint := fmt.Sprintf("%s/api/%s/locations/%d", o.BaseUrl, o.ApiVersion, id)
	request, err := retryablehttp.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Send the request
	response, err := o.Client.Do(request.Request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Parse the response body
	var location domain.Location
	err = ParseResponse(response, &location)
	if err != nil {
		return nil, err
	}

	return &location, nil
}

// CreateJob creates a new job using the Optii SDK.
func (o *OptiiSdk) CreateJob(job *domain.Job) (interface{}, error) {
	jsonBody, err := json.Marshal(job)
	if err != nil {
		return nil, fmt.Errorf("error marshalling job: %w", err)
	}
	requestBody := bytes.NewBuffer(jsonBody)

	// Create a new POST request
	endpoint := fmt.Sprintf("%s/api/%s/jobs", o.BaseUrl, o.ApiVersion)
	request, err := http.NewRequest(http.MethodPost, endpoint, requestBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Send the request
	response, err := o.Client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("client error making request: %w", err)
	}
	defer response.Body.Close()

	// Parse the response body
	result := &domain.JobCreated{}
	err = ParseResponse(response, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetFloorRooms retrieves the list of rooms on a specific floor.
// It takes the floorID as input and returns a slice of domain.Location representing the rooms on the floor.
// If an error occurs during the retrieval process, it returns nil and the error.
func (o *OptiiSdk) GetFloorRooms(floorID int) ([]domain.Location, error) {
	var locations []domain.Location

	first := int32(0)
	next := int32(100)
	for {
		locationsQuery, err := o.GetLocations(first, next, "Room")
		if err != nil {
			return nil, err
		}

		for _, location := range locationsQuery.Locations {
			if location.ParentLocation != nil && location.ParentLocation.ID == floorID {
				locations = append(locations, location)
			}
		}

		if !locationsQuery.PageInfo.HasNextPage {
			break
		}

		next = int32(locationsQuery.PageInfo.EndCursor)
	}

	return locations, nil
}

// GetLocations retrieves a list of locations from the Optii SDK.
// It takes the following parameters:
//   - first: the number of locations to retrieve in the first batch
//   - next: the number of locations to retrieve in subsequent batches
//   - locationType: the type of location to filter by (optional)
//
// The function returns a LocationsQuery object containing the retrieved locations, or an error if the request fails.
func (o *OptiiSdk) GetLocations(first int32, next int32, locationType string) (*LocationsQuery, error) {
	// Create the endpoint
	endpoint := fmt.Sprintf("%s/api/%s/locations?first=%d&next=%d", o.BaseUrl, o.ApiVersion, first, next)
	if locationType != "" {
		endpoint = fmt.Sprintf("%s&locationType=%s", endpoint, locationType)
	}

	// Create a new GET request
	request, err := retryablehttp.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Send the request
	response, err := o.Client.Do(request.Request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Parse the response body
	locationsQuery := &LocationsQuery{}
	err = ParseResponse(response, &locationsQuery)
	if err != nil {
		return nil, err
	}

	return locationsQuery, nil
}

// GetFloorLocations retrieves all locations that belong to a specific floor.
// It takes a floorID as input and returns a slice of domain.Location and an error.
// The function iterates through paginated results of GetLocations and filters the locations
// that have a parent location matching the given floorID.
func (o *OptiiSdk) GetFloorLocations(floorID int) ([]domain.Location, error) {
	var locations []domain.Location

	first := int32(0)
	next := int32(100)
	for {
		locationsQuery, err := o.GetLocations(first, next, "")
		if err != nil {
			return nil, err
		}

		for _, location := range locationsQuery.Locations {
			if location.ParentLocation != nil && location.ParentLocation.ID == floorID {
				locations = append(locations, location)
			}
		}

		if !locationsQuery.PageInfo.HasNextPage {
			break
		}

		next = int32(locationsQuery.PageInfo.EndCursor)
	}

	return locations, nil
}

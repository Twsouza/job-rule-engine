package sdk

import (
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

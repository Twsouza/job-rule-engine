//go:build e2e
// +build e2e

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	port := "3000"

	// Call the main function
	go main()

	// Wait for the server to start
	time.Sleep(3 * time.Second)

	// This payload is a valid one from the Optii API
	payload := `{
		"departmentId": 13,
		"jobItemId": 184,
		"locationsId": [
			2
		]
	}`

	// Send a POST request to the server
	requestBody := bytes.NewBuffer([]byte(payload))
	res, err := http.Post(fmt.Sprintf("http://localhost:%s/v1/jobs", port), "application/json", requestBody)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	jobResult := []domain.JobResult{}
	err = json.NewDecoder(res.Body).Decode(&jobResult)
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.Empty(t, jobResult[0].Err)
}

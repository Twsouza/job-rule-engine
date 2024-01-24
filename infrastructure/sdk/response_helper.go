package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type ErrorResponse struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func ParseResponse(resp *http.Response, data interface{}) error {
	if resp.StatusCode == http.StatusOK || resp.StatusCode < http.StatusCreated {
		err := json.NewDecoder(resp.Body).Decode(data)
		if err != nil {
			return err
		}

		return nil
	}

	/*
		The JSON in the response body is not valid, due to the presence of
		a trailing comma after the last element in the array
		To work around this, we will use a regex to remove the trailing comma
		then parse the JSON using the cleaned body
	*/

	// Use regex to remove trailing commas
	re := regexp.MustCompile(`,\s*}`)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	cleanedBody := re.ReplaceAll(body, []byte("}"))

	errorResponse := &ErrorResponse{}
	err = json.Unmarshal(cleanedBody, errorResponse)
	if err != nil {
		return err
	}

	/* Delete previous code after fix and use this instead
	errorResponse := &ErrorResponse{}
	err := json.NewDecoder(resp.Body).Decode(errorResponse)
	if err != nil {
		return err
	}
	*/

	return fmt.Errorf("%s: %s", errorResponse.Title, errorResponse.Detail)
}

package task

import (
	"errors"
	"regexp"

	"github.com/Twsouza/job-rule-engine/domain"
)

type CleanBedsRoom struct{}

// AssertRule checks if the given job request satisfies the conditions for a clean room task.
// It verifies that the job request belongs to the "Housekeeping" department,
// the job item is either "Blanket", "Sheets", or "Mattress",
// and the location is "Room".
// If all conditions are met, it returns true; otherwise, it returns false.
func (cr *CleanBedsRoom) AssertRule(jobRequest domain.JobRequest) bool {
	if jobRequest.Department != "Housekeeping" {
		return false
	}

	match, err := regexp.MatchString(`(?i)\b(?:Blanket|Sheets|Mattress)\b`, jobRequest.JobItem)
	if err != nil || !match {
		return false
	}

	if len(jobRequest.Locations) == 0 {
		return false
	}

	for _, location := range jobRequest.Locations {
		if location == "Room" {
			return true
		}
	}

	return false
}

// Execute will create a job to clean the bed(s) in the given room
func (cr *CleanBedsRoom) Execute(jobRequest domain.JobRequest) domain.JobResult {
	return domain.JobResult{
		Request: &jobRequest,
		Result:  "",
		Err:     errors.New("not implemented"),
	}
}
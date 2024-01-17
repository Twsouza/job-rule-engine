package task

import (
	"errors"
	"regexp"

	"github.com/Twsouza/job-rule-engine/domain"
)

type CleanBedsRoom struct{}

// AssertRule checks if the given job request satisfies the conditions to clean beds in a room.
// It returns true if the job request meets the following criteria:
// - The department is "Housekeeping"
// - The job item display name contains "Blanket", "Sheets", or "Mattress"
// - At least one location has a location type of "Room"
// Otherwise, it returns false.
func (cr *CleanBedsRoom) AssertRule(jobRequest domain.JobRequest) bool {
	if jobRequest.Department == nil || jobRequest.Department.Name != "Housekeeping" || jobRequest.JobItem == nil {
		return false
	}

	match, err := regexp.MatchString(`(?i)\b(?:Blanket|Sheets|Mattress)\b`, jobRequest.JobItem.DisplayName)
	if err != nil || !match {
		return false
	}

	for _, location := range jobRequest.Locations {
		if location.LocationType != nil && location.LocationType.DisplayName == "Room" {
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

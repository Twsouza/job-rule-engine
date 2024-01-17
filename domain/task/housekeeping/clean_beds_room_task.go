package task

import (
	"errors"
	"regexp"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/domain/task"
)

type CleanBedsRoom struct {
	API task.JobAPI
}

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
	job := domain.Job{
		Action: "clean",
		Department: domain.JDepartment{
			ID: jobRequest.Department.ID,
		},
		Item: domain.JItem{
			Name: jobRequest.JobItem.DisplayName,
		},
	}

	for _, location := range jobRequest.Locations {
		if location.LocationType != nil && location.LocationType.DisplayName == "Room" {
			job.Locations = append(job.Locations, domain.JLocation{
				ID: location.ID,
			})
		}
	}

	if len(job.Locations) == 0 {
		return domain.JobResult{
			Request: &jobRequest,
			Result:  nil,
			Err:     errors.New("no room location found"),
		}
	}

	result, err := cr.API.CreateJob(job)

	return domain.JobResult{
		Request: &jobRequest,
		Result:  result,
		Err:     err,
	}
}

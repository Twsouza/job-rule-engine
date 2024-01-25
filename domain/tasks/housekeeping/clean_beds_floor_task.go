package housekeeping

import (
	"regexp"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/domain/tasks"
)

type CleanBedsFloor struct {
	API tasks.JobAPI
}

// AssertRule checks if the given job request satisfies the conditions for clean the beds in all rooms with a location type of ‘Room’ on that floor.
// It returns true if the job request meets the following criteria:
// - The job request must have a non-nil Department with the name "Housekeeping".
// - The job item display name contains "Blanket", "Sheets", or "Mattress"
// - The job request must have at least one Location with a LocationType that has a display name of "Floor".
// Otherwise, it returns false.
func (cr *CleanBedsFloor) AssertRule(jobRequest domain.JobRequest) bool {
	if jobRequest.Department == nil || jobRequest.Department.Name != "Housekeeping" || jobRequest.JobItem == nil {
		return false
	}

	match, err := regexp.MatchString(`(?i)\b(?:Blanket|Sheets|Mattress)\b`, jobRequest.JobItem.DisplayName)
	if err != nil || !match {
		return false
	}

	for _, location := range jobRequest.Locations {
		if location.LocationType != nil && location.LocationType.DisplayName == "Floor" {
			return true
		}
	}

	return false
}

// Execute will create a job to clean the beds in all rooms with a location type of ‘Room’ on that floor.
func (cr *CleanBedsFloor) Execute(jobRequest domain.JobRequest) domain.JobResult {
	jr := domain.JobResult{
		Request: &jobRequest,
	}

	job := &domain.Job{
		Action: "clean",
		Department: domain.JDepartment{
			ID: jobRequest.Department.ID,
		},
		Item: domain.JItem{
			Name: jobRequest.JobItem.DisplayName,
		},
	}

	locations, err := cr.API.GetFloorRooms(jobRequest.Locations[0].ID)
	if err != nil {
		jr.Err = err.Error()
		return jr
	}

	for _, location := range locations {
		job.Locations = append(job.Locations, domain.JLocation{
			ID: location.ID,
		})
	}

	if len(job.Locations) == 0 {
		jr.Err = "no locations found for this job"
		return jr
	}

	result, err := cr.API.CreateJob(job)
	if err != nil {
		jr.Err = err.Error()
	}
	jr.Result = result

	return jr
}

package roomservice

import (
	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/domain/tasks"
)

type DeliverJobItemLocationTask struct {
	API tasks.JobAPI
}

// AssertRule checks if the given job request satisfies the conditions to create a job to deliver that job item to the given locations.
// The conditions to return true are:
// - The job request must have a non-nil Department and JobItem.
// - The Department name must be "Room Service".
// - The JobItem must have a non-empty DisplayName.
// - The job request must have more than one location.
// If any of these conditions are not met, false is returned.
func (dj *DeliverJobItemLocationTask) AssertRule(jobRequest domain.JobRequest) bool {
	if jobRequest.Department == nil || jobRequest.JobItem == nil {
		return false
	}

	if jobRequest.Department.Name == "Room Service" && jobRequest.JobItem.DisplayName != "" && len(jobRequest.Locations) > 1 {
		return true
	}

	return false
}

// Execute will create a job to deliver that job item to the given locations.
func (dj *DeliverJobItemLocationTask) Execute(jobRequest domain.JobRequest) domain.JobResult {
	job := &domain.Job{
		Action: "deliver",
		Department: domain.JDepartment{
			ID: jobRequest.Department.ID,
		},
		Item: domain.JItem{
			Name: jobRequest.JobItem.DisplayName,
		},
	}

	for _, location := range jobRequest.Locations {
		job.Locations = append(job.Locations, domain.JLocation{
			ID: location.ID,
		})
	}

	result, err := dj.API.CreateJob(job)

	return domain.JobResult{
		Request: &jobRequest,
		Result:  result,
		Err:     err,
	}
}

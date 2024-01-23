package engineering

import (
	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/domain/tasks"
)

type RepairJobItemLocation struct {
	API tasks.JobAPI
}

// AssertRule checks if the given job request meets the criteria for a repair job item at a location.
// It returns true if the job request belongs to the "Engineering" department and has a non-empty job item and at least one location.
// Otherwise, it returns false.
func (rj RepairJobItemLocation) AssertRule(jobRequest domain.JobRequest) bool {
	if jobRequest.Department == nil || jobRequest.JobItem == nil {
		return false
	}

	if jobRequest.Department.Name == "Engineering" && jobRequest.JobItem.DisplayName != "" && len(jobRequest.Locations) > 0 {
		return true
	}

	return false
}

// Execute will create a job to repair the given job item at the given location(s).
func (rj RepairJobItemLocation) Execute(jobRequest domain.JobRequest) domain.JobResult {
	job := domain.Job{
		Action: "repair",
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

	result, err := rj.API.CreateJob(job)

	return domain.JobResult{
		Request: &jobRequest,
		Result:  result,
		Err:     err,
	}
}

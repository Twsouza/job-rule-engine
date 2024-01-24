package engineering

import (
	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/domain/tasks"
)

type RepairJobItemFloor struct {
	API tasks.JobAPI
}

// AssertRule checks if the given job request meets the criteria for a repair job item in all locations on floor task.
// It verifies that the job request belongs to the "Engineering" department
// and has at least one location specified, with "Floor" being the only allowed location.
// It returns true if the job request meets the criteria, otherwise it returns false.
func (rj RepairJobItemFloor) AssertRule(jobRequest domain.JobRequest) bool {
	if jobRequest.Department == nil || jobRequest.JobItem == nil {
		return false
	}

	if jobRequest.Department.Name == "Engineering" && len(jobRequest.Locations) == 1 && jobRequest.Locations[0].LocationType.DisplayName == "Floor" {
		return true
	}

	return false
}

// Execute will create a job to repair the given job item in all locations on that floor.
func (rj RepairJobItemFloor) Execute(jobRequest domain.JobRequest) domain.JobResult {
	jr := domain.JobResult{
		Request: &jobRequest,
	}

	job := &domain.Job{
		Action: "repair",
		Department: domain.JDepartment{
			ID: jobRequest.Department.ID,
		},
		Item: domain.JItem{
			Name: jobRequest.JobItem.DisplayName,
		},
	}

	locations, err := rj.API.GetFloorLocations(jobRequest.Locations[0].ID)
	if err != nil {
		jr.Err = err
		return jr
	}

	for _, location := range locations {
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

package task

import (
	"errors"

	"github.com/Twsouza/job-rule-engine/domain"
)

type RepairJobItemFloor struct{}

// AssertRule checks if the given job request meets the criteria for a repair job item in all locations on floor task.
// It verifies that the job request belongs to the "Engineering" department
// and has at least one location specified, with "Floor" being the only allowed location.
// It returns true if the job request meets the criteria, otherwise it returns false.
func (rj RepairJobItemFloor) AssertRule(jobRequest domain.JobRequest) bool {
	if jobRequest.Department != "Engineering" {
		return false
	}

	if jobRequest.JobItem == "" {
		return false
	}

	if len(jobRequest.Locations) == 1 && jobRequest.Locations[0] == "Floor" {
		return true
	}

	return false
}

func (rj RepairJobItemFloor) Execute(jobRequest domain.JobRequest) domain.JobResult {
	return domain.JobResult{
		Request: &jobRequest,
		Result:  "",
		Err:     errors.New("not implemented"),
	}
}
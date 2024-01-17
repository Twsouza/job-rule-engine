package task

import (
	"errors"

	"github.com/Twsouza/job-rule-engine/domain"
)

type RepairJobItemLocation struct{}

// AssertRule checks if the given job request meets the criteria for a repair job item at a location.
// It returns true if the job request belongs to the "Engineering" department and has a non-empty job item and at least one location.
// Otherwise, it returns false.
func (rj RepairJobItemLocation) AssertRule(jobRequest domain.JobRequest) bool {
	if jobRequest.Department != "Engineering" {
		return false
	}

	if jobRequest.JobItem != "" && len(jobRequest.Locations) > 0 {
		return true
	}

	return false
}

func (rj RepairJobItemLocation) Execute(jobRequest domain.JobRequest) domain.JobResult {
	return domain.JobResult{
		Request: &jobRequest,
		Result:  "",
		Err:     errors.New("not implemented"),
	}
}
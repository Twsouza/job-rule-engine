package task

import (
	"errors"

	"github.com/Twsouza/job-rule-engine/domain"
)

type DeliverJobItemLocationTask struct{}

// AssertRule checks if the given job request satisfies the conditions to create a job to deliver that job item to the given locations.
// The conditions to return true are:
// - The job request must have a non-nil Department and JobItem.
// - The Department name must be "Room Service".
// - The JobItem must have a non-empty DisplayName.
// - The job request must have at least one location.
// If any of these conditions are not met, false is returned.
func (dj *DeliverJobItemLocationTask) AssertRule(jobRequest domain.JobRequest) bool {
	if jobRequest.Department == nil || jobRequest.JobItem == nil {
		return false
	}

	if jobRequest.Department.Name == "Room Service" && jobRequest.JobItem.DisplayName != "" && len(jobRequest.Locations) > 0 {
		return true
	}

	return false
}

func (dj *DeliverJobItemLocationTask) Execute(jobRequest domain.JobRequest) domain.JobResult {
	return domain.JobResult{
		Request: &jobRequest,
		Result:  "",
		Err:     errors.New("not implemented"),
	}
}

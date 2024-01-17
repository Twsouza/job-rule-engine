package task

import (
	"errors"

	"github.com/Twsouza/job-rule-engine/domain"
)

type DeliverJobItemLocationTask struct{}

// AssertRule checks if the given job request satisfies the conditions to create a job to deliver that job item to the given locations.
// AssertRule checks if the given job request satisfies the conditions to deliver an item in all locations.
// 1. Return false if the job request department is not "Room Service".
// 2. Return false if the job item is empty.
// 3. Return true if there are one or more locations specified in the job request.
// 4. Return false if none of the above conditions are satisfied.
func (dj *DeliverJobItemLocationTask) AssertRule(jobRequest domain.JobRequest) bool {
	if jobRequest.Department != "Room Service" {
		return false
	}

	if jobRequest.JobItem == "" {
		return false
	}

	if len(jobRequest.Locations) > 0 {
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

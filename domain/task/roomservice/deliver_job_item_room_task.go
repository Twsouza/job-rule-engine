package task

import (
	"errors"

	"github.com/Twsouza/job-rule-engine/domain"
)

type DeliverJobItemRoomTask struct{}

// AssertRule checks if the given job request satisfies the conditions to deliver an item in all locations.
// - Return false if the job request department is not "Room Service".
// - Return false if the job item is empty.
// - Return true if the job request has only one location and it is "Floor".
// - Return false for all other cases.
func (dj *DeliverJobItemRoomTask) AssertRule(jobRequest domain.JobRequest) bool {
	if jobRequest.Department != "Room Service" {
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

// Execute will create a job to deliver the given job item in all locations with a location type of 'Room' on that floor
func (dj *DeliverJobItemRoomTask) Execute(jobRequest domain.JobRequest) domain.JobResult {
	return domain.JobResult{
		Request: &jobRequest,
		Result:  "",
		Err:     errors.New("not implemented"),
	}
}

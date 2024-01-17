package task

import (
	"errors"

	"github.com/Twsouza/job-rule-engine/domain"
)

type DeliverJobItemRoomTask struct{}

// AssertRule checks if the given job request satisfies the conditions to deliver an item in all locations.
// It returns true if the job request meets the following conditions:
// - The job request has a non-nil Department field with the name "Room Service".
// - The job request has a non-nil JobItem field with a non-empty DisplayName.
// - The job request has exactly one location of type "Floor".
// Otherwise, it returns false.
func (dj *DeliverJobItemRoomTask) AssertRule(jobRequest domain.JobRequest) bool {
	if jobRequest.Department == nil || jobRequest.Department.Name != "Room Service" || jobRequest.JobItem == nil || jobRequest.JobItem.DisplayName == "" {
		return false
	}

	if len(jobRequest.Locations) == 1 && jobRequest.Locations[0].LocationType.DisplayName == "Floor" {
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

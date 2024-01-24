package roomservice

import (
	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/domain/tasks"
)

type DeliverJobItemRoomTask struct {
	API tasks.JobAPI
}

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
	jr := domain.JobResult{
		Request: &jobRequest,
	}

	job := &domain.Job{
		Action: "deliver",
		Department: domain.JDepartment{
			ID: jobRequest.Department.ID,
		},
		Item: domain.JItem{
			Name: jobRequest.JobItem.DisplayName,
		},
	}

	locations, err := dj.API.GetFloorRooms(jobRequest.Locations[0].ID)
	if err != nil {
		jr.Err = err
		return jr
	}

	for _, location := range locations {
		job.Locations = append(job.Locations, domain.JLocation{
			ID: location.ID,
		})
	}

	jr.Result, jr.Err = dj.API.CreateJob(job)

	return jr
}

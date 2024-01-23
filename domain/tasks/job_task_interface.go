package tasks

import "github.com/Twsouza/job-rule-engine/domain"

// JobTask represents a generic task that can be executed.
type JobTask interface {
	// AssertRule checks if the task can be executed based on the given job request.
	AssertRule(jobRequest domain.JobRequest) bool
	// Execute performs the task based on the given job request and returns the result.
	Execute(jobRequest domain.JobRequest) domain.JobResult
}

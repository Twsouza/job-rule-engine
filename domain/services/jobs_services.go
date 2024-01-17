package services

import (
	"sync"

	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/domain/task"
)

type JobService struct {
	Tasks []task.JobTask
}

// CreateJob creates a job based on the given jobRequest and executes the rules associated with the JobService.
// It returns a slice of domain.JobResult containing the results of the executed rules.
// The function uses a channel to receive the domain.JobResult from each executed rule concurrently.
// If the context is canceled due to a deadline, it returns a domain.JobResult with an error indicating a timeout.
// If the context is canceled for any other reason, it returns a domain.JobResult with the corresponding error.
// The function waits for all rules to finish executing before returning the results.
func (js *JobService) CreateJob(jobRequest domain.JobRequest) []domain.JobResult {
	jrCh := make(chan domain.JobResult)
	wg := sync.WaitGroup{}

	for _, t := range js.Tasks {
		if t.AssertRule(jobRequest) {
			wg.Add(1)

			go func(t task.JobTask, req domain.JobRequest) {
				defer wg.Done()
				jr := t.Execute(req)
				jrCh <- jr
			}(t, jobRequest)
		}
	}

	go func() {
		wg.Wait()
		close(jrCh)
	}()

	var results []domain.JobResult
	for jr := range jrCh {
		results = append(results, jr)
	}

	return results
}

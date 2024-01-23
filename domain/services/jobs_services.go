package services

import (
	"sync"

	"github.com/Twsouza/job-rule-engine/application/dto"
	"github.com/Twsouza/job-rule-engine/domain"
	"github.com/Twsouza/job-rule-engine/domain/tasks"
)

type JobService struct {
	Tasks    []tasks.JobTask
	OptiiAPI OptiiApiInterface
}

func NewJobService(tasks []tasks.JobTask) *JobService {
	return &JobService{
		Tasks: tasks,
	}
}

// CreateJob creates a job based on the given jobRequest and executes the rules associated with the JobService.
// It returns a slice of domain.JobResult containing the results of the executed rules.
// The function uses a channel to receive the domain.JobResult from each executed rule concurrently.
// The function waits for all rules to finish executing before returning the results.
func (js *JobService) CreateJob(jobRequest *domain.JobRequest) []domain.JobResult {
	jrCh := make(chan domain.JobResult)
	wg := sync.WaitGroup{}

	for _, t := range js.Tasks {
		// To avoid any rule changing the jobRequest, I'm passing jobRequest as a value to each rule instead of a reference.
		if t.AssertRule(*jobRequest) {
			wg.Add(1)

			go func(t tasks.JobTask, req domain.JobRequest) {
				defer wg.Done()
				jr := t.Execute(req)
				jrCh <- jr
			}(t, *jobRequest)
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

// LoadJob loads a job request by retrieving the department, job item, and locations
// associated with the given JobRequestDto. It uses concurrent goroutines to fetch
// the data and returns the loaded JobRequest along with any errors encountered.
func (js *JobService) LoadJob(reqDto *dto.JobRequestDto) (*domain.JobRequest, []error) {
	var wg sync.WaitGroup
	errChan := make(chan error, 3)
	departmentChan := make(chan *domain.Department, 1)
	jobItemChan := make(chan *domain.JobItem, 1)
	locationsChan := make(chan []domain.Location, 1)

	wg.Add(3)
	go func() {
		defer wg.Done()
		department, err := js.OptiiAPI.GetDepartmentByID(reqDto.DepartmentID)
		if err != nil {
			errChan <- err
			close(departmentChan)
			return
		}
		departmentChan <- department
	}()

	go func() {
		defer wg.Done()
		jobItem, err := js.OptiiAPI.GetJobItemByID(reqDto.JobItemID)
		if err != nil {
			errChan <- err
			close(jobItemChan)
			return
		}
		jobItemChan <- jobItem
	}()

	go func() {
		defer wg.Done()
		locations, err := js.OptiiAPI.GetLocationsByIds(reqDto.LocationsID)
		if err != nil {
			errChan <- err
			close(locationsChan)
			return
		}
		locationsChan <- locations
	}()

	wg.Wait()
	close(errChan)

	errs := []error{}
	for err := range errChan {
		errs = append(errs, err)
	}

	// receiving all values from the channels to avoid memory leaks
	jr := &domain.JobRequest{
		Department: <-departmentChan,
		JobItem:    <-jobItemChan,
		Locations:  <-locationsChan,
	}

	return jr, errs
}

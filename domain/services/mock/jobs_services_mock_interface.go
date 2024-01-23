package mock

import (
	"github.com/Twsouza/job-rule-engine/application/dto"
	"github.com/Twsouza/job-rule-engine/domain"
)

type JobServiceMock struct {
	CreateJobFunc func(jobRequest *domain.JobRequest) []domain.JobResult
	LoadJobFunc   func(dto *dto.JobRequestDto) (*domain.JobRequest, []error)
}

func (m *JobServiceMock) CreateJob(jobRequest *domain.JobRequest) []domain.JobResult {
	return m.CreateJobFunc(jobRequest)
}

func (m *JobServiceMock) LoadJob(dto *dto.JobRequestDto) (*domain.JobRequest, []error) {
	return m.LoadJobFunc(dto)
}

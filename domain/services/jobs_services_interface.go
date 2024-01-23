package services

import (
	"github.com/Twsouza/job-rule-engine/application/dto"
	"github.com/Twsouza/job-rule-engine/domain"
)

type JobServiceInterface interface {
	CreateJob(jobRequest *domain.JobRequest) []domain.JobResult
	LoadJob(dto *dto.JobRequestDto) (*domain.JobRequest, []error)
}

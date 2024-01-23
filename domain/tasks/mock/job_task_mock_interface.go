package mock

import "github.com/Twsouza/job-rule-engine/domain"

type MockRule struct {
	AssertFunc  func(jobRequest domain.JobRequest) bool
	ExecuteFunc func(jobRequest domain.JobRequest) domain.JobResult
}

func (mr *MockRule) AssertRule(jobRequest domain.JobRequest) bool {
	return mr.AssertFunc(jobRequest)
}

func (mr *MockRule) Execute(jobRequest domain.JobRequest) domain.JobResult {
	return mr.ExecuteFunc(jobRequest)
}

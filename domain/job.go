package domain

// JobRequest represents a job creation request
type JobRequest struct {
	Department string   `json:"department"`
	JobItem    string   `json:"job_item"`
	Locations  []string `json:"locations"`
}

type JobResult struct {
	Request *JobRequest `json:"request"`
	Result  interface{} `json:"result"`
	Err     error       `json:"error"`
}

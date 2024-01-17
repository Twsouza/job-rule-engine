package domain

// JobRequest represents a job creation request
type JobRequest struct {
	Department *Department `json:"department"`
	JobItem    *JobItem    `json:"job_item"`
	Locations  []Location  `json:"locations"`
}

type JobResult struct {
	Request *JobRequest `json:"request"`
	Result  interface{} `json:"result"`
	Err     error       `json:"error"`
}

type Job struct {
	Item struct {
		Name string `json:"name"`
	} `json:"item"`
	Department struct {
		ID int `json:"id"`
	} `json:"department"`
	Location []struct {
		ID int `json:"id"`
	} `json:"location"`
	Action string `json:"action"`
}

package domain

type JobRequest struct {
	Department *Department `json:"department"`
	JobItem    *JobItem    `json:"jobItem"`
	Locations  []Location  `json:"locations"`
}

type JobResult struct {
	Request *JobRequest `json:"request"`
	Result  interface{} `json:"result"`
	Err     error       `json:"error"`
}

type Job struct {
	Item       JItem       `json:"item"`
	Department JDepartment `json:"department"`
	Locations  []JLocation `json:"location"`
	Action     string      `json:"action"`
}

type JItem struct {
	Name string `json:"name"`
}

type JDepartment struct {
	ID int `json:"id"`
}

type JLocation struct {
	ID int `json:"id"`
}

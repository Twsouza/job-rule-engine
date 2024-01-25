package domain

import "time"

type JobRequest struct {
	Department *Department `json:"department"`
	JobItem    *JobItem    `json:"jobItem"`
	Locations  []Location  `json:"locations"`
}

type JobResult struct {
	Request *JobRequest `json:"request"`
	Result  interface{} `json:"result"`
	Err     string      `json:"error"`
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

type JobCreated struct {
	ID   int `json:"id"`
	Item struct {
		Displayname string `json:"displayname"`
	} `json:"item"`
	Type        string `json:"type"`
	Priority    string `json:"priority"`
	Action      string `json:"action"`
	Attachments any    `json:"attachments"`
	Locations   []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"locations"`
	Departments []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"departments"`
	Roles []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"roles"`
	Notes    any       `json:"notes"`
	Assignee any       `json:"assignee"`
	DueBy    time.Time `json:"dueBy"`
}

package sdk

import "github.com/Twsouza/job-rule-engine/domain"

type ErrorResponse struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

type PageInfo struct {
	EndCursor   int  `json:"endCursor"`
	HasNextPage bool `json:"hasNextPage"`
	TotalCount  int  `json:"totalCount"`
}

type LocationsQuery struct {
	PageInfo  *PageInfo         `json:"pageInfo"`
	Locations []domain.Location `json:"items"`
}

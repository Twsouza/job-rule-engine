package dto

type JobRequestDto struct {
	DepartmentID int64   `json:"departmentId"`
	JobItemID    int64   `json:"jobItemId"`
	LocationsID  []int64 `json:"locationsId"`
}

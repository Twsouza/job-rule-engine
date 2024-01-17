package domain

type LocationType struct {
	ID          int    `json:"id"`
	DisplayName string `json:"displayName"`
}

type Location struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	DisplayName    string `json:"displayName"`
	ParentLocation struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
	} `json:"parentLocation"`
	LocationType *LocationType `json:"locationType"`
}

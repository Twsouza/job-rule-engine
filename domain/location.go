package domain

type LocationType struct {
	ID          int    `json:"id"`
	DisplayName string `json:"displayName"`
}

// ParentLocation is created separately from Location to avoid circular reference.
type ParentLocation struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type Location struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	DisplayName    string          `json:"displayName"`
	ParentLocation *ParentLocation `json:"parentLocation,omitempty"`
	LocationType   *LocationType   `json:"locationType,omitempty"`
}

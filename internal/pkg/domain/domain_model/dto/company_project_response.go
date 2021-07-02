package dto

// CompanyProjectsResponse struct
type CompanyProjectsResponse struct {
	Name     string    `json:"name"`
	Projects []Project `json:"projects"`
}

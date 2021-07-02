package dto

// UserProjectsResponse struct
type UserProjectsResponse struct {
	Projects []Project
}

// Project struct
type Project struct {
	Name              string `json:"name"`
	Category          string `json:"category"`
	ProjectedSpend    int    `json:"projected_spend"`
	ProjectedVariance int    `json:"projected_variance"`
	RevenueRecognised int    `json:"revenue_recognised"`
}

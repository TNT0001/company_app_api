package dto

// CreatePojectResponse struct
type CreateProjectResponse struct {
	Name              string `json:"name"`
	Category          string `json:"category"`
	ProjectedSpend    int    `json:"projected_spend,omitempty"`
	ProjectedVariance int    `json:"projected_variance,omitempty"`
	RevenueRecognised int    `json:"revenue_recognised,omitempty"`
}

package dto

// CreateProjectRequest struct
type CreateProjectRequest struct {
	Name              string `form:"name" json:"name" binding:"required,max=255"`
	Category          string `form:"category" json:"category" binding:"required,oneof='client' 'non-billable' 'system'"`
	ProjectedSpend    int    `form:"projected_spend" json:"projected_spend" binding:"omitempty"`
	ProjectedVariance int    `form:"projected_variance" json:"projected_variance"`
	RevenueRecognised int    `form:"revenue_recognised" json:"revenue_recognised"`
}

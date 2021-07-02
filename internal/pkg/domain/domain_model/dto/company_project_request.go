package dto

// CompanyProjectsRequest struct
type CompanyProjectsRequest struct {
	Name string `form:"name" binding:"required,max=255"`
}

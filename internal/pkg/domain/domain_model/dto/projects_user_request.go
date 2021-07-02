package dto

// UserProjectsRequest struct
type UserProjectsRequest struct {
	UserID string `form:"user_id" binding:"required"`
}

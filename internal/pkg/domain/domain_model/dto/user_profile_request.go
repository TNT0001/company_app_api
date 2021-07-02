package dto

// UserProfileRequest struct
type UserProfileRequest struct {
	UID string `form:"uid" binding:""`
}

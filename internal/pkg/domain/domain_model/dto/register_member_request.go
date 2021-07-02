package dto

// RegisterMemberRequest struct
type RegisterMemberRequest struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

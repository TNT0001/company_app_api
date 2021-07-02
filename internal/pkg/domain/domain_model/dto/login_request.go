package dto

// LoginRequest struct
type LoginRequest struct {
	ID       string `form:"id" binding:"required"`
	Password string `form:"password" binding:"required"`
}

package dto

// UserUpdateRequest struct
type UserUpdateRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,password,min=6,max=10"`
	Username string `form:"username" binding:"required,max=16"`
	Birthday string `form:"birthday" binding:"omitempty,timeFormat"`
	ImageUrl string `form:"image_url" binding:"omitempty,urlcustom"`
}

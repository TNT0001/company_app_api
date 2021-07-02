package dto

// UserUpdateResponse struct
type UserUpdateResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Birthday string `json:"birthday,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}

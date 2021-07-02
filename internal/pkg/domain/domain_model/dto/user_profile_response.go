package dto

// UserProfileResponse struct
type UserProfileResponse struct {
	UID  string `json:"uid"`
	User User   `json:"user"`
}

// User struct
type User struct {
	Img             string    `json:"img,omitempty"`
	Name            string    `json:"name,omitempty"`
	UseStartDate    string    `json:"use_start_date,omitempty"`
	Email           string    `json:"email,omitempty"`
	Birthday        string    `json:"birthday,omitempty"`
	LearningPurpose int       `json:"learning_purpose,omitempty"`
	Projects        []Project `json:"projects"`
}

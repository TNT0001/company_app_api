package model

import "time"

// BaseModel struct
type BaseModel struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:datetime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime;not null"`
}

// BaseModelWithDeleted struct
type BaseModelWithDeleted struct {
	BaseModel
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;type:datetime;null"`
}

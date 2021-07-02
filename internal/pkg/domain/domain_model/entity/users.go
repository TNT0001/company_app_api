package entity

import (
	"go-api/pkg/shared/gorm/model"
	"time"
)

// UsersTableName TableName
var UsersTableName = "users"

// Users struct
type Users struct {
	ID             int        `gorm:"column:id;primary_key;type:int(11);not null"`
	Username       string     `gorm:"column:username;not null"`
	Email          string     `gorm:"column:email;not null"`
	Password       string     `gorm:"column:password;not null"`
	Birthday       *time.Time `gorm:"column:birthday"`
	ImageURL       *string    `gorm:"column:image_url;type:varchar(2083)"`
	RefreshToken   *string    `gorm:"column:refresh_token"`
	Token          *string    `gorm:"column:token"`
	TokenExpriedAt *time.Time `gorm:"column:token_expried_at"`
	IsActive       bool       `gorm:"column:is_active"`
	CompanyID      int        `gorm:"column:company_id"`
	model.BaseModel
}

// TableName func
func (i *Users) TableName() string {
	return UsersTableName
}

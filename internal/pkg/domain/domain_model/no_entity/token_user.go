package noentity

import (
	"time"
)

// TokenUser struct
type TokenUser struct {
	ID             int       `gorm:"column:id;primary_key;type:int(11);not null"`
	UserID         int       `gorm:"column:user_id"`
	Token          string    `gorm:"column:token"`
	TokenExpriedAt time.Time `gorm:"column:token_expried_at"`
}

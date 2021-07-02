package entity

import (
	"go-api/pkg/shared/gorm/model"
)

// ProjectsTableName TableName
var ProjectsUserTableName = "projects_user"

// Users struct
type ProjectsUser struct {
	ID        int `gorm:"column:id;primary_key;type:int(11);not null"`
	ProjectID int `gorm:"column:project_id;not null;type:int(11) unsigned"`
	UserID    int `gorm:"column:user_id;not null;type:int(11) unsigned"`
	model.BaseModel
}

// TableName func
func (p *ProjectsUser) TableName() string {
	return ProjectsUserTableName
}

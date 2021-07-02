package entity

import (
	"go-api/pkg/shared/gorm/model"
)

// ProjectsTableName TableName
var ProjectsTableName = "projects"

// Users struct
type Projects struct {
	ID                int    `gorm:"column:id;primary_key;type:int(11) unsigned;not null"`
	Name              string `gorm:"column:name;not null"`
	Category          string `gorm:"column:category;not null;enum('client','non-billable','system');default:'client'"`
	ProjectedSpend    int    `gorm:"column:projected_spend;not null;default:0"`
	ProjectedVariance int    `gorm:"column:projected_variance;not null;default:0"`
	RevenueRecognised int    `gorm:"column:revenue_recognised;not null;default:0"`
	model.BaseModel
}

// TableName func
func (p *Projects) TableName() string {
	return ProjectsTableName
}

package entity

import "go-api/pkg/shared/gorm/model"

const CompanyTableName = `companies`

type Companies struct {
	ID   int    `gorm:"column:id;type:int(11);primary key;auto_increment;not null"`
	Name string `gorm:"column:name;type:varchar(255);not null"`
	model.BaseModel
}

func (i *Companies) TableName() string {
	return CompanyTableName
}

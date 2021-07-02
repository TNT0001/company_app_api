package service

import (
	"go-api/internal/pkg/domain/domain_model/entity"
)

// ProjectsRepository interface
type ProjectsRepository interface {
	FindProjectByUserID(userID int) ([]entity.Projects, error)
	CreateProject(user entity.Users, project entity.Projects) error
}

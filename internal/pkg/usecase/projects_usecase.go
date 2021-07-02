package usecase

import (
	"go-api/internal/pkg/domain/domain_model/dto"
	"go-api/internal/pkg/domain/domain_model/entity"
	"go-api/internal/pkg/domain/service"
)

// ProjectsInterface interface
type ProjectsInterface interface {
	GetProjectsByUserID(userID int) (dto.UserProjectsResponse, error)
	PostCreateProject(user entity.Users, req dto.CreateProjectRequest) (dto.CreateProjectResponse, error)
}

// UsersUsecase struct
type ProjectsUsecase struct {
	*BaseUsecase
	repo service.ProjectsRepository
}

// NewProjectUsecase func
func NewProjectUsecase(bu *BaseUsecase, pr service.ProjectsRepository) ProjectsInterface {
	return &ProjectsUsecase{BaseUsecase: bu, repo: pr}
}

// GetProjectsByUserID func
func (pu *ProjectsUsecase) GetProjectsByUserID(userID int) (dto.UserProjectsResponse, error) {
	// Get list entity projects
	projects, err := pu.repo.FindProjectByUserID(userID)
	if err != nil {
		return dto.UserProjectsResponse{}, err
	}

	//create projects response
	ProjectsResponse := dto.UserProjectsResponse{}
	for _, project_entity := range projects {
		project := dto.Project{
			Name:              project_entity.Name,
			Category:          project_entity.Category,
			ProjectedSpend:    project_entity.ProjectedSpend,
			ProjectedVariance: project_entity.ProjectedVariance,
			RevenueRecognised: project_entity.RevenueRecognised,
		}
		ProjectsResponse.Projects = append(ProjectsResponse.Projects, project)
	}

	return ProjectsResponse, nil
}

// CreateProject func
func (pu *ProjectsUsecase) PostCreateProject(user entity.Users, req dto.CreateProjectRequest) (dto.CreateProjectResponse, error) {
	project := entity.Projects{
		Name:              req.Name,
		Category:          req.Category,
		ProjectedSpend:    req.ProjectedSpend,
		ProjectedVariance: req.ProjectedVariance,
		RevenueRecognised: req.RevenueRecognised,
	}

	err := pu.repo.CreateProject(user, project)
	if err != nil {
		return dto.CreateProjectResponse{}, err
	}

	return dto.CreateProjectResponse(req), nil
}

package usecase

import (
	"errors"
	"go-api/internal/pkg/domain/domain_model/dto"
	"go-api/internal/pkg/domain/service"
	"go-api/pkg/infrastructure"

	"github.com/jinzhu/gorm"
)

// CompaniesUsecaseInterface type
type CompaniesUsecaseInterface interface {
	GetProjects(req dto.CompanyProjectsRequest) (dto.CompanyProjectsResponse, error)
}

// CopaniesUsecase struct
type companiesUsecase struct {
	BaseUsecase
	repo service.CompaniesRepositoryInterface
}

// NewCompaniesUsecase func
func NewCompaniesUsecase(bu *BaseUsecase, repo service.CompaniesRepositoryInterface) CompaniesUsecaseInterface {
	return &companiesUsecase{
		BaseUsecase: *bu,
		repo:        repo,
	}
}

// GetGetProjects func
func (cu *companiesUsecase) GetProjects(req dto.CompanyProjectsRequest) (dto.CompanyProjectsResponse, error) {
	companyName := req.Name

	_, err := cu.repo.GetCompanyByName(companyName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.CompanyProjectsResponse{}, errors.New(infrastructure.ErrCompanyNotFound)
		}
		return dto.CompanyProjectsResponse{}, err
	}

	projects, err := cu.repo.GetProjects(companyName)
	if err != nil {
		return dto.CompanyProjectsResponse{}, err
	}

	response := dto.CompanyProjectsResponse{}
	response.Name = companyName

	for _, project := range projects {
		response.Projects = append(response.Projects, dto.Project{
			Name:              project.Name,
			Category:          project.Category,
			ProjectedSpend:    project.ProjectedSpend,
			ProjectedVariance: project.ProjectedVariance,
			RevenueRecognised: project.RevenueRecognised,
		})
	}

	return response, nil
}

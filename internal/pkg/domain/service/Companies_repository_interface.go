package service

import "go-api/internal/pkg/domain/domain_model/entity"

//CompaniesRepositoryInterface type
type CompaniesRepositoryInterface interface {
	GetProjects(name string) ([]entity.Projects, error)
	GetCompanyByName(name string) (entity.Companies, error)
}

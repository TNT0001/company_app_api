package repository

import (
	"go-api/internal/pkg/domain/domain_model/entity"
	"go-api/internal/pkg/domain/service"
	"go-api/pkg/infrastructure"
)

// CompaniesRepository struct
type companiesRepository struct {
	BaseRepository
	DB infrastructure.Database
}

// NewCompanyRepository func
func NewCompanyRepository(br *BaseRepository, db infrastructure.Database) service.CompaniesRepositoryInterface {
	return &companiesRepository{
		BaseRepository: *br,
		DB:             db,
	}
}

// GetProjects func
func (c *companiesRepository) GetProjects(name string) ([]entity.Projects, error) {
	// create projects slice
	projects := make([]entity.Projects, 0)

	// sql query
	query := `SELECT DISTINCT projects.*
	FROM projects
	INNER JOIN projects_user
	ON projects.id = projects_user.project_id
	WHERE projects_user.user_id IN (SELECT id
	FROM users
	WHERE users.company_id = (SELECT id
	FROM companies
	WHERE companies.name = ?));`

	err := c.DB.Query(&projects, query, name)
	if err != nil {
		return []entity.Projects{}, nil
	}

	return projects, nil
}

// GetCompanyByName func
func (c *companiesRepository) GetCompanyByName(name string) (entity.Companies, error) {
	company := entity.Companies{}
	company.Name = name
	err := c.DB.Find(&company, &company)
	if err != nil {
		return entity.Companies{}, err
	}
	return company, nil
}

package repository

import (
	"go-api/internal/pkg/domain/domain_model/entity"
	"go-api/internal/pkg/domain/service"
	"go-api/pkg/infrastructure"
)

// projectsRepository struct
type projectsRepository struct {
	BaseRepository
	DB infrastructure.Database
}

// NewProjectsRepository func
func NewProjectsRepository(br *BaseRepository, db infrastructure.Database) service.ProjectsRepository {
	return &projectsRepository{BaseRepository: *br, DB: db}
}

// FindProjectByUserID func
func (p *projectsRepository) FindProjectByUserID(userID int) ([]entity.Projects, error) {
	projects := make([]entity.Projects, 0)

	// check if id is not in database
	err := p.DB.Find(userID, &entity.Users{})
	if err != nil {
		return nil, err
	}
	// Query projects statement
	query := `SELECT * FROM projects JOIN projects_user ON projects.id = projects_user.project_id WHERE projects_user.user_id = ?;`

	err = p.DB.Query(&projects, query, userID)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

// CreateProject func
func (p *projectsRepository) CreateProject(user entity.Users, project entity.Projects) error {
	// Begin transaction
	tx := p.DB.Begin()

	// Insert project into db
	if err := p.DB.CreateWithTransaction(tx, &project); err != nil {
		p.DB.Rollback(tx)
		return err
	}

	// Create projects_user entity
	projectsUser := entity.ProjectsUser{
		UserID:    user.ID,
		ProjectID: project.ID,
	}

	// Insert projectsUser into db
	if err := p.DB.CreateWithTransaction(tx, &projectsUser); err != nil {
		p.DB.Rollback(tx)
		return err
	}

	// Commit transaction
	return p.DB.Commit(tx)
}

package repository

import (
	"go-api/internal/pkg/domain/domain_model/entity"
	noentity "go-api/internal/pkg/domain/domain_model/no_entity"
	"go-api/internal/pkg/domain/service"
	"go-api/pkg/infrastructure"
)

// usersRepository struct
type usersRepository struct {
	BaseRepository
	DB infrastructure.Database
}

// VerifyUserToken struct
type VerifyUserToken struct {
	entity.Users
	noentity.TokenUser
}

// NewUsersRepository func
func NewUsersRepository(br *BaseRepository, db infrastructure.Database) service.UsersRepository {
	return &usersRepository{BaseRepository: *br, DB: db}
}

// FindUser func
func (u *usersRepository) FindUser(condition entity.Users) (entity.Users, error) {
	user := entity.Users{}
	err := u.DB.Find(condition, &user)
	if u.DB.IsRecordNotFoundError(err) {
		return user, err
	}
	return user, nil
}

// CreateUser func
func (u *usersRepository) CreateUser(user entity.Users) (entity.Users, error) {
	err := u.DB.Create(&user)
	if err != nil {
		return entity.Users{}, nil
	}
	return user, nil
}

// UpdateUserTokenLogin func
func (u *usersRepository) UpdateUserTokenLogin(user, oldUser entity.Users) error {
	err := u.DB.Update(false, entity.Users{}, &oldUser, &user)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserProfile func
func (u *usersRepository) UpdateUserProfile(user, oldUser entity.Users) error {
	err := u.DB.Update(false, entity.Users{}, &oldUser, &user)
	if err != nil {
		return err
	}
	return nil
}

// GetUserProjects func
func (u *usersRepository) GetUserProjects(user entity.Users) ([]entity.Projects, error) {
	projects := make([]entity.Projects, 0)

	query := `SELECT * FROM projects JOIN projects_user ON projects.id = projects_user.project_id WHERE projects_user.user_id = ?;`

	err := u.DB.Query(&projects, query, user.ID)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

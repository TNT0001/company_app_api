package service

import (
	"go-api/internal/pkg/domain/domain_model/entity"
)

// UsersRepository interface
type UsersRepository interface {
	FindUser(condition entity.Users) (entity.Users, error)
	CreateUser(user entity.Users) (entity.Users, error)
	UpdateUserTokenLogin(user, oldUser entity.Users) error
	UpdateUserProfile(user, oldUser entity.Users) error
	GetUserProjects(user entity.Users) ([]entity.Projects, error)
}

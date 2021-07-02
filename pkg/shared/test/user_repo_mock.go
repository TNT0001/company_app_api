package test

import (
	"go-api/internal/pkg/domain/domain_model/entity"

	"github.com/stretchr/testify/mock"
)

// UserRepoMock struct
type UserRepoMock struct {
	mock.Mock
}

// FindUser func
func (r *UserRepoMock) FindUser(condition entity.Users) (entity.Users, error) {
	args := r.Called(condition)
	return args.Get(0).(entity.Users), args.Error(1)
}

// CreateUser func
func (r *UserRepoMock) CreateUser(user entity.Users) (entity.Users, error) {
	args := r.Called(user)
	return args.Get(0).(entity.Users), args.Error(1)
}

// UpdateUserTokenLogin func
func (r *UserRepoMock) UpdateUserTokenLogin(user, oldUser entity.Users) error {
	args := r.Called(user, oldUser)
	return args.Error(0)
}

// UpdateUserProfile func
func (r *UserRepoMock) UpdateUserProfile(user, oldUser entity.Users) error {
	args := r.Called(user, oldUser)
	return args.Error(0)
}

// GetUserProjects func
func (r *UserRepoMock) GetUserProjects(user entity.Users) ([]entity.Projects, error) {
	args := r.Called(user)
	return args.Get(0).([]entity.Projects), args.Error(1)
}

package test

import (
	"go-api/internal/pkg/domain/domain_model/entity"

	"github.com/stretchr/testify/mock"
)

// Project repository mock struct
type ProjectRepoMock struct {
	mock.Mock
}

// FIndProjectByUserID func
func (p *ProjectRepoMock) FindProjectByUserID(userID int) ([]entity.Projects, error) {
	args := p.Called(userID)
	return args.Get(0).([]entity.Projects), args.Error(1)
}

// CreateProject func
func (p *ProjectRepoMock) CreateProject(user entity.Users, project entity.Projects) error {
	args := p.Called(user, project)
	return args.Error(0)
}

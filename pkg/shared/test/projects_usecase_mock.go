package test

import (
	"go-api/internal/pkg/domain/domain_model/dto"
	"go-api/internal/pkg/domain/domain_model/entity"

	"github.com/stretchr/testify/mock"
)

// ProjectUsecase mock obj
type ProjectUsecaseMock struct {
	mock.Mock
}

// GetProjectsByUserID func
func (p *ProjectUsecaseMock) GetProjectsByUserID(userID int) (dto.UserProjectsResponse, error) {
	args := p.Called(userID)
	return args.Get(0).(dto.UserProjectsResponse), args.Error(1)
}

// PostCreateProject func
func (p *ProjectUsecaseMock) PostCreateProject(user entity.Users, req dto.CreateProjectRequest) (dto.CreateProjectResponse, error) {
	args := p.Called(user, req)
	return args.Get(0).(dto.CreateProjectResponse), args.Error(1)
}

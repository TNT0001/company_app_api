package usecase

import (
	"errors"
	"go-api/internal/pkg/domain/domain_model/dto"
	"go-api/internal/pkg/domain/domain_model/entity"
	"go-api/pkg/infrastructure"
	"go-api/pkg/shared/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestProjectsCreateUsecaseSuccess func
func TestProjectsCreateUsecaseSuccess(t *testing.T) {
	user := entity.Users{}

	req := dto.CreateProjectRequest{
		Name:              "test1",
		Category:          "client",
		ProjectedSpend:    0,
		ProjectedVariance: 0,
		RevenueRecognised: 0,
	}

	repo := new(test.ProjectRepoMock)
	ProjectUsecase := NewProjectUsecase(bu, repo)

	t.Run("test create project usecase success", func(t *testing.T) {
		repo.On("CreateProject", user, mock.Anything).Return(nil)
		_, err := ProjectUsecase.PostCreateProject(user, req)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

// TestProjectsCreateUsecaseFail func
func TestProjectsCreateUsecaseFail(t *testing.T) {
	user := entity.Users{}

	req := dto.CreateProjectRequest{
		Name:              "test1",
		Category:          "client",
		ProjectedSpend:    0,
		ProjectedVariance: 0,
		RevenueRecognised: 0,
	}

	repo := new(test.ProjectRepoMock)
	ProjectUsecase := NewProjectUsecase(bu, repo)

	t.Run("test create project usecase success", func(t *testing.T) {
		repo.On("CreateProject", user, mock.Anything).Return(errors.New(infrastructure.ErrRecordNotFound))
		_, err := ProjectUsecase.PostCreateProject(user, req)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

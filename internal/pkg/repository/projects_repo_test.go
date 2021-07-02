package repository

import (
	"errors"
	"go-api/internal/pkg/domain/domain_model/entity"
	"go-api/pkg/infrastructure"
	"go-api/pkg/shared/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProjectSuccess(t *testing.T) {
	user := entity.Users{
		ID:       1,
		Email:    "lannt@gmail.com",
		Password: "12345467",
	}
	project := entity.Projects{
		ID:       1,
		Name:     "prj1",
		Category: "client",
	}

	t.Run("test create project success", func(t *testing.T) {
		db := new(test.DatabaseMock)
		tx := new(test.DatabaseMock)
		db.On("Begin").Return(tx)
		db.On("CreateWithTransaction", mock.Anything, &project).Return(nil)
		db.On("CreateWithTransaction", mock.Anything, mock.Anything).Return(nil)
		db.On("Commit", mock.Anything).Return(nil)
		pr := NewProjectsRepository(br, db)
		err := pr.CreateProject(user, project)
		assert.Nil(t, err)
		db.AssertExpectations(t)
	})
}

func TestCreateProjectFail(t *testing.T) {
	user := entity.Users{
		ID:       1,
		Email:    "lannt@gmail.com",
		Password: "12345467",
	}
	project := entity.Projects{
		ID:       1,
		Name:     "prj1",
		Category: "client",
	}

	t.Run("test create project fail", func(t *testing.T) {
		db := new(test.DatabaseMock)
		tx := new(test.DatabaseMock)
		db.On("Begin").Return(tx)
		db.On("CreateWithTransaction", mock.Anything, &project).Return(errors.New(infrastructure.ErrRecordNotFound))
		db.On("Rollback", mock.Anything).Return(nil)
		pr := NewProjectsRepository(br, db)
		err := pr.CreateProject(user, project)
		assert.NotNil(t, err)
		db.AssertExpectations(t)
	})
}

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

var userID = "user ID"
var logger = new(test.LoggerMock)
var br = NewBaseRepository(logger)

// TestFindUserSuccess func
func TestFindUserSuccess(t *testing.T) {
	condition := entity.Users{
		Email:    "lannt@gmail.com",
		Password: "12345467",
	}

	t.Run("test find user sucess", func(t *testing.T) {
		db := new(test.DatabaseMock)
		db.On("Find", condition, mock.Anything).Return(nil)
		db.On("IsRecordNotFoundError", mock.Anything).Return(false)
		ar := NewUsersRepository(br, db)
		_, err := ar.FindUser(condition)
		assert.Nil(t, err)
		db.AssertExpectations(t)
	})
}

// TestFindUserFail func
func TestFindUserFail(t *testing.T) {
	condition := entity.Users{
		Email:    "lannt@gmail.com",
		Password: "12345467",
	}

	t.Run("test find user fail", func(t *testing.T) {
		db := new(test.DatabaseMock)
		db.On("Find", condition, mock.Anything).Return(errors.New(infrastructure.ErrRecordNotFound))
		db.On("IsRecordNotFoundError", mock.Anything).Return(true)
		ar := NewUsersRepository(br, db)
		_, err := ar.FindUser(condition)
		assert.NotNil(t, err)
		db.AssertExpectations(t)
	})
}

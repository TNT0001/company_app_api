package usecase

import (
	"errors"
	"go-api/internal/pkg/domain/domain_model/dto"
	"go-api/internal/pkg/domain/domain_model/entity"
	"go-api/pkg/infrastructure"
	"go-api/pkg/shared/test"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userID = "user ID"
var logger = new(test.LoggerMock)
var bu = NewBaseUsecase(logger)

// TestGetAccountSuccess func
func TestGetUserTokenLoginSuccess(t *testing.T) {
	tests := []dto.LoginRequest{
		{
			ID:       "lannt@gmail.com",
			Password: "12345467",
		},
		{
			ID:       "lannt1@gmail.com",
			Password: "1234546",
		},
	}

	for i, v := range tests {
		t.Run("test get user login "+strconv.Itoa(i+1), func(t *testing.T) {
			ur := new(test.UserRepoMock)
			ur.On("FindUser", mock.Anything).Return(entity.Users{Password: v.Password, Email: v.ID}, nil)
			ur.On("UpdateUserTokenLogin", mock.Anything, mock.Anything).Return(nil)
			uu := NewUserUseCase(bu, ur)
			_, err := uu.GetUserTokenLogin(v)
			assert.Nil(t, err)
			ur.AssertExpectations(t)
		})
	}
}

func TestGetUserTokenLoginFail(t *testing.T) {
	req := dto.LoginRequest{
		ID:       "lannt1@gmail.com",
		Password: "12345467",
	}
	ur := new(test.UserRepoMock)
	ur.On("FindUser", mock.Anything).Return(entity.Users{}, errors.New(infrastructure.ErrRecordNotFound))
	uu := NewUserUseCase(bu, ur)
	_, err := uu.GetUserTokenLogin(req)
	assert.NotNil(t, err)
	ur.AssertExpectations(t)
}

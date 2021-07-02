package test

import (
	"go-api/internal/pkg/domain/domain_model/dto"
	"go-api/internal/pkg/domain/domain_model/entity"

	"github.com/stretchr/testify/mock"
)

// UserUsecaseMock struct
type UserUsecaseMock struct {
	mock.Mock
}

// GetUserTokenLogin func
func (u *UserUsecaseMock) GetUserTokenLogin(req dto.LoginRequest) (string, error) {
	args := u.Called(req)
	return args.Get(0).(string), args.Error(1)
}

// PostCreateUser func
func (u *UserUsecaseMock) PostCreateUser(req dto.RegisterMemberRequest) (dto.RegisterMemberResponse, error) {
	args := u.Called(req)
	return args.Get(0).(dto.RegisterMemberResponse), args.Error(1)
}

// GetUserProfile func
func (u *UserUsecaseMock) GetUserProfile(user entity.Users) (dto.User, error) {
	args := u.Called(user)
	return args.Get(0).(dto.User), args.Error(1)
}

// PatchUpdateUser func
func (u *UserUsecaseMock) PatchUpdateUser(req dto.UserUpdateRequest, oldUser entity.Users) (dto.UserUpdateResponse, error) {
	args := u.Called(req, oldUser)
	return args.Get(0).(dto.UserUpdateResponse), args.Error(1)
}

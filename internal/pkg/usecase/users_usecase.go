package usecase

import (
	"errors"
	"go-api/internal/pkg/domain/domain_model/dto"
	"go-api/internal/pkg/domain/domain_model/entity"
	"go-api/internal/pkg/domain/service"
	"go-api/pkg/infrastructure"
	"go-api/pkg/shared/auth"
	"go-api/pkg/shared/utils"
	"time"

	uuid "github.com/satori/go.uuid"
)

// UsersInterface interface
type UsersInterface interface {
	PostCreateUser(req dto.RegisterMemberRequest) (dto.RegisterMemberResponse, error)
	GetUserTokenLogin(req dto.LoginRequest) (string, error)
	GetUserProfile(user entity.Users) (dto.User, error)
	PatchUpdateUser(req dto.UserUpdateRequest, oldUser entity.Users) (dto.UserUpdateResponse, error)
}

// UsersUsecase struct
type UsersUsecase struct {
	BaseUsecase
	repo service.UsersRepository
}

// GetUserTokenLogin func
func (u *UsersUsecase) GetUserTokenLogin(req dto.LoginRequest) (string, error) {
	user, err := u.repo.FindUser(entity.Users{
		Email:    req.ID,
		Password: utils.EncryptPassword(req.Password),
	})

	if err != nil {
		return "", err
	}

	if !user.IsActive {
		return "", errors.New(infrastructure.ErrLoginFail)
	}

	timeNow := time.Now()

	timeExpriedAt := timeNow.Add(utils.TimeExpriedDuration)

	// generate uuid
	uuid := uuid.Must(uuid.NewV4(), nil)
	tokenString, err := auth.GenerateJWTToken(auth.JWTParam{
		UUID:       uuid,
		Authorized: true,
		ExpriedAt:  timeExpriedAt,
	})
	if err != nil {
		return "", err
	}

	newUser := entity.Users{
		Token:          &tokenString,
		TokenExpriedAt: &timeExpriedAt,
	}

	err = u.repo.UpdateUserTokenLogin(newUser, user)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// PostCreateUser func
func (u *UsersUsecase) PostCreateUser(req dto.RegisterMemberRequest) (dto.RegisterMemberResponse, error) {
	user, err := u.repo.FindUser(entity.Users{
		Email: req.Email,
	})

	if err != nil {
		return dto.RegisterMemberResponse{}, err
	}

	if (user.Email != "") && user.IsActive {
		return dto.RegisterMemberResponse{}, errors.New(infrastructure.ErrEmailAlreadyExist)
	}

	if (user.Email != "") && !user.IsActive {
		return dto.RegisterMemberResponse{}, errors.New(infrastructure.ErrEmailAuthentication)
	}

	user = entity.Users{
		Email:    req.Email,
		Password: utils.EncryptPassword(req.Password),
		IsActive: true,
	}
	_, err = u.repo.CreateUser(user)
	if err != nil {
		return dto.RegisterMemberResponse{}, err
	}
	return dto.RegisterMemberResponse{}, nil
}

// GetUserProfile func
func (u *UsersUsecase) GetUserProfile(user entity.Users) (dto.User, error) {
	// profile data
	profile := dto.User{
		Name:  user.Username,
		Email: user.Email,
	}

	if user.ImageURL != nil {
		profile.Img = *user.ImageURL
	}

	if user.Birthday != nil {
		profile.Birthday = utils.GetDateToString(user.Birthday)
	}

	// get projects
	projects := make([]dto.Project, 0)
	result, err := u.repo.GetUserProjects(user)
	if err != nil {
		return dto.User{}, err
	}
	for _, item := range result {
		projects = append(projects, dto.Project{
			Name:              item.Name,
			Category:          item.Category,
			ProjectedSpend:    item.ProjectedSpend,
			ProjectedVariance: item.ProjectedVariance,
			RevenueRecognised: item.RevenueRecognised,
		})
	}

	profile.Projects = projects
	return profile, nil
}

// PatchUpdateUser func
func (u *UsersUsecase) PatchUpdateUser(req dto.UserUpdateRequest, oldUser entity.Users) (dto.UserUpdateResponse, error) {
	// check req.Email already exist and active
	user, err := u.repo.FindUser(entity.Users{
		Email: req.Email,
	})
	if err != nil {
		return dto.UserUpdateResponse{}, err
	}
	if user.Email != "" && user.IsActive {
		return dto.UserUpdateResponse{}, errors.New(infrastructure.ErrEmailAlreadyExist)
	}

	// check duplicate pass
	newPassword := utils.EncryptPassword(req.Password)
	if newPassword == oldUser.Password {
		return dto.UserUpdateResponse{}, errors.New(infrastructure.ErrPasswordInvalid)
	}

	// newUser variable hold update information
	newUser := entity.Users{
		Email:    req.Email,
		Password: newPassword,
		Username: req.Username,
	}

	// if req.Birthday and req.ImageUrl not zero value pass them to newUser
	if req.Birthday != "" {
		birthday, err := time.Parse(utils.FormatTime, req.Birthday)
		if err != nil {
			return dto.UserUpdateResponse{}, errors.New(infrastructure.ErrBirthday)
		}
		newUser.Birthday = &birthday
	}
	if req.ImageUrl != "" {
		newUser.ImageURL = &req.ImageUrl
	}

	// update user profile
	err = u.repo.UpdateUserProfile(newUser, oldUser)
	if err != nil {
		return dto.UserUpdateResponse{}, err
	}

	return dto.UserUpdateResponse(req), nil
}

// NewUserUseCase response new Usecase instance
func NewUserUseCase(bu *BaseUsecase, r service.UsersRepository) *UsersUsecase {
	return &UsersUsecase{BaseUsecase: *bu, repo: r}
}

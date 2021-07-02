package users

import (
	"go-api/internal/pkg/domain/domain_model/dto"
	"go-api/internal/pkg/repository"
	"go-api/internal/pkg/usecase"
	"go-api/pkg/infrastructure"
	"go-api/pkg/shared/handler"
	"go-api/pkg/shared/middleware"
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
)

// HTTPHandler struct
type HTTPHandler struct {
	handler.ApplicationHTTPHandler
	usecase usecase.UsersInterface
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(ah *handler.ApplicationHTTPHandler, r *repository.BaseRepository, u *usecase.BaseUsecase, db infrastructure.Database) *HTTPHandler {
	usersRepository := repository.NewUsersRepository(r, db)
	usersUsecase := usecase.NewUserUseCase(u, usersRepository)
	return &HTTPHandler{ApplicationHTTPHandler: *ah, usecase: usersUsecase}
}

// Login func
func (h *HTTPHandler) Login(c *gin.Context) {
	req := dto.LoginRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error: &dto.ErrorResponse{
				ErrorMessage: err.Error(),
			},
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}

	token, err := h.usecase.GetUserTokenLogin(req)
	if err != nil {
		code := http.StatusInternalServerError
		if strings.Contains(err.Error(), infrastructure.ErrRecordNotFound) || strings.Contains(err.Error(), infrastructure.ErrLoginFail) {
			code = http.StatusBadRequest
		}
		data := dto.BaseResponse{
			Status: code,
			Error: &dto.ErrorResponse{
				ErrorMessage: infrastructure.ErrLoginFail,
			},
		}
		c.JSON(code, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: dto.LoginResponse{
			UID:   "1KoYg02dqOTQ6KM1oNkDC8Db5ZK2",
			Token: token,
		},
	}
	c.JSON(http.StatusOK, data)
	return
}

// RegisterMember func
func (h *HTTPHandler) RegisterMember(c *gin.Context) {
	req := dto.RegisterMemberRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error: &dto.ErrorResponse{
				ErrorMessage: err.Error(),
			},
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	_, err = h.usecase.PostCreateUser(req)
	if err != nil {
		code := http.StatusInternalServerError
		if strings.Contains(err.Error(), infrastructure.ErrEmailAlreadyExist) || strings.Contains(err.Error(), infrastructure.ErrEmailAuthentication) {
			code = http.StatusBadRequest
		}
		data := dto.BaseResponse{
			Status: code,
			Error: &dto.ErrorResponse{
				ErrorMessage: err.Error(),
			},
		}
		c.JSON(code, data)
		return
	}

	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: dto.RegisterMemberResponse{
			UID: "1KoYg02dqOTQ6KM1oNkDC8Db5ZK2",
		},
	}

	c.JSON(http.StatusOK, data)
}

// GetUserProfile func
func (h *HTTPHandler) GetUserProfile(c *gin.Context) {
	req := dto.UserProfileRequest{}

	err := c.ShouldBindQuery(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error: &dto.ErrorResponse{
				ErrorMessage: err.Error(),
			},
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}

	user := middleware.GetUserFromContext(c)

	userProfile, err := h.usecase.GetUserProfile(user)

	if err != nil {
		code := http.StatusInternalServerError
		if strings.Contains(err.Error(), infrastructure.ErrRecordNotFound) {
			code = http.StatusBadRequest
		}
		data := dto.BaseResponse{
			Status: code,
			Error: &dto.ErrorResponse{
				ErrorMessage: err.Error(),
			},
		}
		c.JSON(code, data)
		return
	}

	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: dto.UserProfileResponse{
			UID:  req.UID,
			User: userProfile,
		},
	}

	c.JSON(http.StatusOK, data)
}

// UpdateUserProfile func
func (h *HTTPHandler) UpdateUserProfile(c *gin.Context) {
	// binding request
	req := dto.UserUpdateRequest{}

	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error: &dto.ErrorResponse{
				ErrorMessage: err.Error(),
			},
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}

	oldUser := middleware.GetUserFromContext(c)

	// update user
	response, err := h.usecase.PatchUpdateUser(req, oldUser)

	if err != nil {
		code := http.StatusInternalServerError
		if strings.Contains(err.Error(), infrastructure.ErrEmailAlreadyExist) ||
			strings.Contains(err.Error(), infrastructure.ErrPasswordInvalid) {
			code = http.StatusBadRequest
		}
		data := dto.BaseResponse{
			Status: code,
			Error: &dto.ErrorResponse{
				ErrorMessage: err.Error(),
			},
		}
		c.JSON(code, data)
		return
	}

	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: response,
	}
	c.JSON(http.StatusOK, data)
}

package projects

import (
	"go-api/internal/pkg/domain/domain_model/dto"
	"go-api/internal/pkg/repository"
	"go-api/internal/pkg/usecase"
	"go-api/pkg/infrastructure"
	"go-api/pkg/shared/handler"
	"go-api/pkg/shared/middleware"
	"go-api/pkg/shared/utils"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// ProjectHandler struct
type ProjectHandler struct {
	handler.ApplicationHTTPHandler
	projectUsecase usecase.ProjectsInterface
}

// NewProjectHandler func
func NewProjectHandler(ah *handler.ApplicationHTTPHandler, r *repository.BaseRepository, u *usecase.BaseUsecase, db infrastructure.Database) *ProjectHandler {
	pr := repository.NewProjectsRepository(r, db)
	pu := usecase.NewProjectUsecase(u, pr)
	return &ProjectHandler{*ah, pu}
}

// GetProjects func
func (ph *ProjectHandler) GetProjects(c *gin.Context) {
	// get user_id param
	userID, ok := c.Params.Get("user_id")
	if !ok {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error: &dto.ErrorResponse{
				ErrorMessage: "missing user_id",
			},
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}

	// validate user_id
	ID, err := strconv.Atoi(userID)
	if err != nil || ID < 1 {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error: &dto.ErrorResponse{
				ErrorMessage: infrastructure.ErrInvalidID,
			},
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}

	// get projects
	projects, err := ph.projectUsecase.GetProjectsByUserID(ID)
	if err != nil {
		code := http.StatusInternalServerError
		if strings.Contains(err.Error(), infrastructure.ErrRecordNotFound) {
			code = http.StatusNotFound
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
	// get csv file
	filename, err := utils.CSVFromPojects(projects)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusInternalServerError,
			Error: &dto.ErrorResponse{
				ErrorMessage: err.Error(),
			},
		}
		c.JSON(http.StatusInternalServerError, data)
		return
	}
	defer os.Remove(filename)

	c.File(filename)
}

// CreateProject func
func (ph *ProjectHandler) CreateProject(c *gin.Context) {
	req := dto.CreateProjectRequest{}
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

	user := middleware.GetUserFromContext(c)

	project, err := ph.projectUsecase.PostCreateProject(user, req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusInternalServerError,
			Error: &dto.ErrorResponse{
				ErrorMessage: err.Error(),
			},
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}

	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: project,
	}
	c.JSON(http.StatusOK, data)
}

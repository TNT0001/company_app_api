package companies

import (
	"errors"
	"go-api/internal/pkg/domain/domain_model/dto"
	"go-api/internal/pkg/repository"
	"go-api/internal/pkg/usecase"
	"go-api/pkg/infrastructure"
	"go-api/pkg/shared/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

// companiesHandler struct
type companiesHandler struct {
	handler.ApplicationHTTPHandler
	companiesUsecase usecase.CompaniesUsecaseInterface
}

// NewCompaniesHandler func
func NewCompaniesHandler(ah *handler.ApplicationHTTPHandler, r *repository.BaseRepository, u *usecase.BaseUsecase, db infrastructure.Database) *companiesHandler {
	repo := repository.NewCompanyRepository(r, db)
	us := usecase.NewCompaniesUsecase(u, repo)
	return &companiesHandler{
		ApplicationHTTPHandler: *ah,
		companiesUsecase:       us,
	}
}

// GetCompanyProjects func
func (u *companiesHandler) GetCompanyProjects(c *gin.Context) {
	req := dto.CompanyProjectsRequest{}
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

	projects, err := u.companiesUsecase.GetProjects(req)
	if err != nil {
		code := http.StatusBadRequest
		if errors.Is(err, errors.New(infrastructure.ErrCompanyNotFound)) {
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

	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: projects,
	}
	c.JSON(http.StatusOK, data)
}

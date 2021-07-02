package projects

import (
	"encoding/json"
	"fmt"
	"go-api/internal/pkg/domain/domain_model/dto"
	"go-api/internal/pkg/domain/domain_model/entity"
	"go-api/internal/pkg/repository"
	"go-api/internal/pkg/usecase"
	"go-api/pkg/shared/handler"
	"go-api/pkg/shared/test"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	// Function
	CreateFn = "PostCreateProject"

	//Path
	CreatePath = "/api/app/projects/create"

	// Method
	GET  = "GET"
	POST = "POST"

	// Info
	LogInfoFn = "Info"
	// Error
	LogErrFn = "Error"
	//Field
)

func TestNewProjectHandlerSuccess(t *testing.T) {
	db := new(test.DatabaseMock)
	loger := new(test.LoggerMock)
	ah := handler.NewApplicationHTTPHandler(loger)
	bu := usecase.NewBaseUsecase(loger)
	r := repository.NewBaseRepository(loger)
	ph := NewProjectHandler(ah, r, bu, db)
	assert.IsType(t, &ProjectHandler{}, ph)
}

func TestCreateProjectSuccess(t *testing.T) {
	// create project handler
	logger := new(test.LoggerMock)
	logger.On(LogErrFn, mock.Anything, mock.Anything).Return()
	ah := handler.NewApplicationHTTPHandler(logger)
	pu := new(test.ProjectUsecaseMock)
	pu.On(CreateFn, mock.Anything, mock.Anything).Return(dto.CreateProjectResponse{
		Name:              "pr1",
		Category:          "client",
		ProjectedSpend:    0,
		ProjectedVariance: 0,
		RevenueRecognised: 0,
	}, nil)
	ph := &ProjectHandler{
		ApplicationHTTPHandler: *ah,
		projectUsecase:         pu,
	}

	// create request body contains json data, create httptest.record
	req := dto.CreateProjectRequest{
		Name:     "pr1",
		Category: "client",
	}
	json, err := json.Marshal(&req)
	if err != nil {
		panic(err)
	}
	fakerBody := strings.NewReader(string(json))
	fmt.Println(fakerBody)
	request := httptest.NewRequest(POST, CreatePath, fakerBody)
	request.Header.Set("content-type", "application/json")
	w := httptest.NewRecorder()

	// gin router, middlerware handler ...
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("user", entity.Users{
			ID:       1,
			Username: "tung",
			Password: "tungbn42123",
		})
	})
	router.POST("/api/app/projects/create", ph.CreateProject)
	router.ServeHTTP(w, request)
	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body)
}

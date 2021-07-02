package users

import (
	"encoding/json"
	"fmt"
	"go-api/internal/pkg/domain/domain_model/dto"
	"go-api/internal/pkg/repository"
	"go-api/internal/pkg/usecase"
	"go-api/pkg/shared/handler"
	"go-api/pkg/shared/test"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	// Function
	LoginFn = "Login"

	//Path
	LoginPath = "/api/app/login"

	// Method
	GET  = "GET"
	POST = "POST"

	// Info
	LogInfoFn = "Info"
	// Error
	LogErrFn = "Error"
	//Field
	accountIDField = "user_id"
)

func TestNewHTTPHandler(t *testing.T) {
	logger := new(test.LoggerMock)
	ah := handler.NewApplicationHTTPHandler(logger)
	database := new(test.DatabaseMock)
	br := repository.NewBaseRepository(logger)
	bu := usecase.NewBaseUsecase(logger)
	ach := NewHTTPHandler(ah, br, bu, database)
	assert.IsType(t, &HTTPHandler{}, ach)
}

// TestLoginSuccess func
func TestLoginSuccess(t *testing.T) {

	body := dto.LoginRequest{
		ID:       "lannt@gmail.com",
		Password: "12345",
	}
	stringBody, _ := json.Marshal(body)
	fakerBody := strings.NewReader(string(stringBody))
	fmt.Println(fakerBody)
	logger := new(test.LoggerMock)
	logger.On(LogErrFn, mock.Anything, mock.Anything).Return()
	ah := handler.NewApplicationHTTPHandler(logger)
	uu := new(test.UserUsecaseMock)
	ach := HTTPHandler{*ah, uu}
	uu.On(LoginFn, mock.Anything).Return(dto.LoginResponse{}, nil)
	r := gin.Default()

	r.POST(LoginPath, ach.Login)
	w := httptest.NewRecorder()
	request, err := http.NewRequest(POST, LoginPath, fakerBody)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(request)
	r.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body)

}

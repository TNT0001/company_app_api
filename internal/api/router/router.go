package router

import (
	"go-api/internal/pkg/domain/domain_model/dto"
	"go-api/internal/pkg/handler/companies"
	"go-api/internal/pkg/handler/projects"
	"go-api/internal/pkg/handler/users"
	"go-api/internal/pkg/repository"
	"go-api/internal/pkg/usecase"
	"go-api/pkg/infrastructure"
	"go-api/pkg/shared/handler"
	"go-api/pkg/shared/middleware"
	"go-api/pkg/shared/utils"
	"go-api/pkg/shared/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router is application struct
type Router struct {
	Engine        *gin.Engine
	DBHandler     infrastructure.Database
	LoggerHandler infrastructure.Logger
}

// InitializeRouter initializes Engine and middleware
func (r *Router) InitializeRouter() {
	r.Engine.Use(gin.Logger())
	r.Engine.Use(gin.Recovery())
}

// SetupHandler set database and redis and usecase.
func (r *Router) SetupHandler() {
	// base handle
	ah := handler.NewApplicationHTTPHandler(r.LoggerHandler)
	// base usecase
	bu := usecase.NewBaseUsecase(r.LoggerHandler)
	// base repository
	br := repository.NewBaseRepository(r.LoggerHandler)
	// user handler
	ur := users.NewHTTPHandler(ah, br, bu, r.DBHandler)
	// project handler
	ph := projects.NewProjectHandler(ah, br, bu, r.DBHandler)
	_ = validator.New()
	// company handler
	ch := companies.NewCompaniesHandler(ah, br, bu, r.DBHandler)

	// health check
	r.Engine.GET("/health_check", func(c *gin.Context) {
		data := dto.BaseResponse{
			Status: http.StatusOK,
			Result: gin.H{"message": "Health check OK!"},
		}
		c.JSON(http.StatusOK, data)
	})

	if utils.GetStringFlag("env") == "dev" || utils.GetStringFlag("env") == "local" {
		r.Engine.Static("/api-spec", utils.GetStringFlag("workdir")+"api")
		r.Engine.Static("/swagger", utils.GetStringFlag("workdir")+"third_party/swagger_ui")
	}

	// router api
	api := r.Engine.Group("/api")
	{
		// router api for app
		appAPI := api.Group("/app")
		{
			appAPI.POST("/login", ur.Login)
			appAPI.Use(middleware.AuthMiddleware(*br, r.DBHandler))
			{
				// user API
				userAPI := appAPI.Group("/user")
				{
					// get user profile
					userAPI.GET("/profile", ur.GetUserProfile)

					// update user route
					userAPI.PATCH("/update", ur.UpdateUserProfile)
				}

				// project API
				projectAPI := appAPI.Group("/projects")
				{
					// get projects by user id
					projectAPI.GET("/:user_id", ph.GetProjects)

					// create project
					projectAPI.POST("/create", ph.CreateProject)
				}

				// company api
				companyAPI := appAPI.Group("/company")
				{
					// get company projects
					companyAPI.GET("/projects", ch.GetCompanyProjects)
				}
			}
		}

		// router api for web
		webAPI := api.Group("/web")
		{
			webAPI.POST("/login", ur.Login)
			webAPI.POST("/register_member", ur.RegisterMember)
		}
	}

}

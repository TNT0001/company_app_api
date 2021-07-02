package middleware

import (
	"go-api/internal/pkg/domain/domain_model/dto"
	"go-api/internal/pkg/domain/domain_model/entity"
	"go-api/internal/pkg/repository"
	"go-api/pkg/infrastructure"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	// AtJwtKey is used to create the Access token signature
	AtJwtKey = []byte("access secret")
)

const (
	// JWTHeaderKey header key jwt token
	JWTHeaderKey = "Authorization"

	// PrefixAuthorization prefix token
	PrefixAuthorization = "Bearer "

	// UserKey string
	UserKey = "user"

	// InvalidToken token is invalid
	InvalidToken = "Invalid Token"
)

// GetUserFromContext get user from context
func GetUserFromContext(c *gin.Context) entity.Users {
	value, exist := c.Get(UserKey)
	if !exist {
		return entity.Users{}
	}
	return value.(entity.Users)
}

// AuthMiddleware func
func AuthMiddleware(baseRepo repository.BaseRepository, db infrastructure.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.GetHeader(JWTHeaderKey)
		if clientToken == "" {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error: &dto.ErrorResponse{
					ErrorMessage: "Authorization Token is required",
				},
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}

		claims := jwt.MapClaims{}

		extractedToken := strings.Split(clientToken, PrefixAuthorization)

		if (len(extractedToken) == 2) && (extractedToken[0] == "") {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error: &dto.ErrorResponse{
					ErrorMessage: InvalidToken,
				},
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}

		// Parse the claims
		parsedToken, err := jwt.ParseWithClaims(clientToken, claims, func(token *jwt.Token) (interface{}, error) {
			return AtJwtKey, nil
		})

		if err != nil {
			if err.Error() == jwt.ErrSignatureInvalid.Error() {
				data := dto.BaseResponse{
					Status: http.StatusUnauthorized,
					Error: &dto.ErrorResponse{
						ErrorMessage: InvalidToken,
					},
				}
				c.JSON(http.StatusUnauthorized, data)
				c.Abort()
				return
			}
			// token is expried
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error: &dto.ErrorResponse{
					ErrorMessage: err.Error(),
				},
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}

		// check user exist in DB
		userRepo := repository.NewUsersRepository(&baseRepo, db)
		user, err := userRepo.FindUser(entity.Users{
			Token: &clientToken,
		})

		if err != nil {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error: &dto.ErrorResponse{
					ErrorMessage: err.Error(),
				},
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}

		if user == (entity.Users{}) {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error: &dto.ErrorResponse{
					ErrorMessage: InvalidToken,
				},
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}

		// set value for UserKey
		c.Set(UserKey, user)

		if !parsedToken.Valid {
			data := dto.BaseResponse{
				Status: http.StatusUnauthorized,
				Error: &dto.ErrorResponse{
					ErrorMessage: InvalidToken,
				},
			}
			c.JSON(http.StatusUnauthorized, data)
			c.Abort()
			return
		}
		c.Next()
	}

}

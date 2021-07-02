package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// JWTParam struct
type JWTParam struct {
	UUID       uuid.UUID
	Authorized bool
	ExpriedAt  time.Time
}

const (
	// ParseTokenErr parse token error message
	ParseTokenErr = "Parse token error"
)

// ContextKey string type
type ContextKey string

// GenerateJWTToken generate jwt token
func GenerateJWTToken(params JWTParam) (string, error) {
	var goClaims = jwt.MapClaims{}
	goClaims["authorized"] = params.Authorized
	goClaims["uuid"] = params.UUID
	goClaims["exp"] = params.ExpriedAt.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, goClaims)

	tokenString, err := token.SignedString([]byte("access secret"))
	return tokenString, err
}

// ParseJWTTokenWithClaims parse jwt token
func ParseJWTTokenWithClaims(tokenStr string) (data map[string]interface{}, err error) {
	if token, _ := jwt.Parse(tokenStr, nil); token != nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			return claims, claims.Valid()
		}
	}
	return nil, err
}

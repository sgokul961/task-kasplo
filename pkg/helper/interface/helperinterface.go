package interfaceshelper

import (
	"main.go/pkg/models"
)

type Helper interface {
	GenerateToken(admin models.Users) (string, error)
	CheckPasswordHash(password string, hash string) bool
	ChekEmailFormat(email string) bool
	VerifyPassword(requestPassword, dbPassword string) bool
	HashPassword(password string) string
	//GenerateRefreshToken(userid int64, email string, key string) (string, error)
	//ValidateToken(tokenString string) (*jwt.Token, error)
}

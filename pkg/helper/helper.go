package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"main.go/pkg/config"
	interfaceshelper "main.go/pkg/helper/interface"
	"main.go/pkg/models"
)

type authCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
type helper struct {
	cfg config.Config
}

func NewHelper(config config.Config) interfaceshelper.Helper {
	return &helper{cfg: config}
}

func (h *helper) GenerateToken(user models.Login) (string, error) {

	claims := &authCustomClaims{
		Email: user.Email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("newcode"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// func (h *helper) PasswordHashing(password string) (string, error) {

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

// 	if err != nil {
// 		return "", errors.New("internal server error")
// 	}

// 	hash := string(hashedPassword)
// 	return hash, nil

// }

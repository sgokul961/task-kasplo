package helper

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"main.go/pkg/config"
	interfaceshelper "main.go/pkg/helper/interface"
	"main.go/pkg/models"
)

type authCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}
type helper struct {
	cfg config.Config
}

func NewHelper(config config.Config) interfaceshelper.Helper {
	return &helper{cfg: config}
}

func (h *helper) GenerateToken(user models.Users) (string, error) {

	claims := &authCustomClaims{
		Email: user.Email,
		Id:    user.UserID,
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

func (l *helper) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func (l *helper) VerifyPassword(requestPassword, dbPassword string) bool {
	requestPassword = fmt.Sprintf("%x", md5.Sum([]byte(requestPassword)))
	fmt.Println("req", requestPassword)
	return requestPassword == dbPassword
}

func (l *helper) HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "cant hash"
	}
	return string(bytes)
}

func (l *helper) ChekEmailFormat(email string) bool {
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex
	re := regexp.MustCompile(emailRegexPattern)

	// Validate the email using the regex
	return re.MatchString(email)

}

// func (h *helper) ValidateToken(tokenString string) (*jwt.Token, error) {
// 	//var con config.Config
// 	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
// 		}
// 		return []byte("newcode"), nil
// 	})
// 	if err != nil {
// 		// There was an error during token validation
// 		fmt.Println("Error validating token:", err)
// 	}

// 	if claims, ok := token.Claims.(*authCustomClaims); ok && token.Valid {
// 		fmt.Printf("Token ID: %d\n", claims.Id)
// 		fmt.Printf("Token Role: %s\n", claims.Email)
// 		// Add more logging or inspection of claims as needed
// 	} else {
// 		fmt.Println("Invalid token claims")
// 	}

// 	return token, err
// }

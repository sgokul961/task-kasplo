package repo

import (
	"fmt"

	"gorm.io/gorm"
	"main.go/pkg/models"
	interfaces "main.go/pkg/repo/interface"
)

type LoginRepo struct {
	DB *gorm.DB
}

func NewLoginRepo(DB *gorm.DB) interfaces.Loginrepo {
	return &LoginRepo{
		DB: DB,
	}
}
func (l *LoginRepo) Login(logindetails models.Login) (models.Login, error) {
	fmt.Println("logg", logindetails.Email)
	var login models.Users

	if err := l.DB.Raw(`SELECT * FROM users WHERE email=$1 `, logindetails.Email).Scan(&login).Error; err != nil {
		return models.Login{}, err
	}
	fmt.Println("user_ aher ", login)
	return models.Login{
		Email:    login.Email,
		Password: login.Password,
	}, nil

}
func (l *LoginRepo) Signup(signup models.Users) error {

	err := l.DB.Exec(`INSERT INTO users (name,email,password) values (?,?,?)`, signup.Name, signup.Email, signup.Password).Error

	if err != nil {
		return err
	}
	return nil

}

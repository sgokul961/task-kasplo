package interfaces

import "main.go/pkg/models"

type Loginrepo interface {
	Login(logindetails models.Login) (models.Login, error)
	Signup(signup models.Users) error
}

package interfaceshelper

import "main.go/pkg/models"

type Helper interface {
	GenerateToken(admin models.Login) (string, error)
	//PasswordHashing(password string) (string, error)
}

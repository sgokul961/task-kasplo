package usecaseinterface

import (
	"main.go/pkg/models"
	"main.go/pkg/utils"
)

type LOginUseCase interface {
	Login(loginDetails models.Login) (utils.TokenLOgin, error)
	Signup(signupdeatils models.Users) error
}

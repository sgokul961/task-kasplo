package usecaseinterface

import (
	"main.go/pkg/demodb"
	"main.go/pkg/models"
	"main.go/pkg/utils"
)

type UseCase interface {
	Login(loginDetails models.Users) (utils.TokenLOgin, error)
	Signup(signupdeatils models.Users) (models.SignupRes, error)
	AddToDo(add demodb.ToDo) (demodb.ToDo, error)
	UpdateToDO(update demodb.UpdateToDo, todoID int, userId int) (demodb.UpdateToDo, error)
	DeleteToDO(todoId int, userId int) (string, error)
	// GetToDoIdOfUser()
}

package interfaces

import (
	"main.go/pkg/demodb"
	"main.go/pkg/models"
)

type Loginrepo interface {
	Login(logindetails models.Users) (models.Users, error)
	Signup(signup models.Users) (models.SignupRes, error)
	CheckUserExistance(email string) bool

	AddToDo(content demodb.ToDo) (demodb.ToDo, error)
	UpdateToDo(content demodb.UpdateToDo) (demodb.UpdateToDo, error)
	DeleteToDo(id int) (string, error)

	//supporting functions
	CheToDOExist(id int) bool
	GetToDoByID(id int) (demodb.ToDo, error)
	FetchUserEmail(id int) (string, error)
	CheckUserIdExist(id int) bool

	//GetToDoIdOfUser()

}

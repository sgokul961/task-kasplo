package usecase

import (
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"main.go/pkg/demodb"
	interfaceshelper "main.go/pkg/helper/interface"
	"main.go/pkg/models"
	interfaces "main.go/pkg/repo/interface"
	"main.go/pkg/usecase/usecaseinterface"
	"main.go/pkg/utils"
)

type UseCase struct {
	Loginrepo interfaces.Loginrepo
	helper    interfaceshelper.Helper
}

func NewUseCase(repo interfaces.Loginrepo, help interfaceshelper.Helper) usecaseinterface.UseCase {
	return &UseCase{
		Loginrepo: repo,
		helper:    help,
	}
}

func (l *UseCase) Signup(signupdeatils models.Users) (models.SignupRes, error) {
	exist := l.Loginrepo.CheckUserExistance(signupdeatils.Email)

	if exist {
		return models.SignupRes{}, errors.New("user already exists")
	}

	hashePassword := l.helper.HashPassword(signupdeatils.Password)

	if hashePassword == "" {
		return models.SignupRes{}, errors.New("failed to hash password")
	}

	signupdeatils.Password = hashePassword

	emailFormat := l.helper.ChekEmailFormat(signupdeatils.Email)

	if !emailFormat {
		return models.SignupRes{}, errors.New("email format is wrong")

	}

	res, err := l.Loginrepo.Signup(signupdeatils)

	if err != nil {
		return models.SignupRes{}, err
	}

	return res, nil
}

func (l *UseCase) Login(loginDetails models.Users) (utils.TokenLOgin, error) {

	// exist := l.Loginrepo.CheckUserExistance(loginDetails.Email)

	// if !exist {
	// 	return utils.TokenLOgin{}, errors.New("User dosent exist")
	// }

	login_detail, err := l.Loginrepo.Login(loginDetails)

	if err != nil {
		return utils.TokenLOgin{}, err

	}
	err = bcrypt.CompareHashAndPassword([]byte(login_detail.Password), []byte(loginDetails.Password))

	if err != nil {
		return utils.TokenLOgin{}, errors.New("wrong password")
	}

	var logindetailsResponse models.Users

	err = copier.Copy(&logindetailsResponse, &login_detail)

	if err != nil {
		return utils.TokenLOgin{}, err
	}
	tokenString, err := l.helper.GenerateToken(logindetailsResponse)
	fmt.Println("token", tokenString)

	if err != nil {
		return utils.TokenLOgin{}, err
	}

	return utils.TokenLOgin{
		LoginResponse: utils.LoginResponse{
			UserID: logindetailsResponse.UserID,
			Name:   logindetailsResponse.Name,
			Email:  logindetailsResponse.Email,
		},
		Token: tokenString,
	}, nil

}
func (l *UseCase) AddToDo(add demodb.ToDo) (demodb.ToDo, error) {

	idExist := l.Loginrepo.CheckUserIdExist(add.UserID)
	if !idExist {
		return demodb.ToDo{}, errors.New("user ID does not exist")
	}
	mailId, err := l.Loginrepo.FetchUserEmail(add.UserID)
	if err != nil {
		return demodb.ToDo{}, err
	}

	UserExist := l.Loginrepo.CheckUserExistance(mailId)

	if !UserExist {
		return demodb.ToDo{}, err
	}
	if add.Title == "" || add.Description == "" || add.DueDate.IsZero() {
		return demodb.ToDo{}, errors.New("invalid input data")
	}

	create, err := l.Loginrepo.AddToDo(add)
	if err != nil {
		return demodb.ToDo{}, err

	}
	return create, nil

}

func (l *UseCase) UpdateToDO(update demodb.UpdateToDo, todoId int, userId int) (demodb.UpdateToDo, error) {

	currentToDo, err := l.Loginrepo.GetToDoByID(todoId)

	if err != nil {
		return demodb.UpdateToDo{}, errors.New("ToDo ID does not exist")
	}

	if currentToDo.UserID != userId {
		return demodb.UpdateToDo{}, errors.New("unauthorized access: you are not the owner of this ToDo item")

	}
	if update.Title == "" || update.Description == "" || update.DueDate.IsZero() {
		return demodb.UpdateToDo{}, errors.New("invalid input data")

	}

	update.ID = uint(todoId)

	updatetodo, err := l.Loginrepo.UpdateToDo(update)

	if err != nil {
		return demodb.UpdateToDo{}, err
	}
	return updatetodo, nil

}
func (l *UseCase) DeleteToDO(todoId int, userId int) (string, error) {
	exist := l.Loginrepo.CheToDOExist(todoId)
	if !exist {
		return "", errors.New("Id doesnot exist")
	}
	// Check if the ToDo ID exists and get the current ToDo item
	currentToDo, err := l.Loginrepo.GetToDoByID(todoId)
	if err != nil {
		return "", errors.New("ToDo ID does not exist")
	}

	// Verify that the authenticated user is the owner of the ToDo item
	if currentToDo.UserID != userId {
		return "", errors.New("unauthorized access: you are not the owner of this ToDo item")
	}

	// Perform the delete
	result, err := l.Loginrepo.DeleteToDo(todoId)
	if err != nil {
		return "", err
	}

	return result, nil
}

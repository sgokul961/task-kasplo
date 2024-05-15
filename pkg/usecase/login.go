package usecase

import (
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	interfaceshelper "main.go/pkg/helper/interface"
	"main.go/pkg/models"
	interfaces "main.go/pkg/repo/interface"
	"main.go/pkg/usecase/usecaseinterface"
	"main.go/pkg/utils"
)

type LOginUseCase struct {
	Loginrepo interfaces.Loginrepo
	helper    interfaceshelper.Helper
}

func NewLoginUseCase(repo interfaces.Loginrepo, help interfaceshelper.Helper) usecaseinterface.LOginUseCase {
	return &LOginUseCase{
		Loginrepo: repo,
		helper:    help,
	}
}
func (l *LOginUseCase) Login(loginDetails models.Login) (utils.TokenLOgin, error) {
	logindetail, err := l.Loginrepo.Login(loginDetails)

	if err != nil {
		return utils.TokenLOgin{}, err
	}
	fmt.Println("user: ", loginDetails, " user: ", logindetail)
	if loginDetails.Password != logindetail.Password {
		return utils.TokenLOgin{}, errors.New("wrong password")
	}
	// err = bcrypt.CompareHashAndPassword([]byte(loginDetails.Password), []byte(loginDetails.Password))
	// fmt.Println(err)
	// if err != nil {
	// 	return utils.TokenLOgin{}, err

	var logindetailsResponse models.Login

	err = copier.Copy(&logindetailsResponse, &logindetail)

	if err != nil {
		return utils.TokenLOgin{}, err
	}
	tokenString, err := l.helper.GenerateToken(logindetailsResponse)

	if err != nil {
		return utils.TokenLOgin{}, err
	}

	return utils.TokenLOgin{
		Admin: logindetailsResponse,
		Token: tokenString,
	}, nil

}
func (l *LOginUseCase) Signup(signupdeatils models.Users) error {
	//hashePassword, err := l.helper.PasswordHashing(signupdeatils.Password)

	// if err != nil {
	// 	return errors.New("ErrorHashingPassword")

	// }
	// signupdeatils.Password = hashePassword
	err := l.Loginrepo.Signup(signupdeatils)

	if err != nil {
		return err
	}

	return nil
}

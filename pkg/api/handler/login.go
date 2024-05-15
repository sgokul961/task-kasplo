package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/pkg/models"
	"main.go/pkg/usecase/usecaseinterface"
	"main.go/pkg/utils/response"
)

type LoginHNadler struct {
	loginusecase usecaseinterface.LOginUseCase
}

func NewLoginHandler(usecase usecaseinterface.LOginUseCase) *LoginHNadler {
	return &LoginHNadler{
		loginusecase: usecase,
	}
}
func (l *LoginHNadler) Login(c *gin.Context) {
	var logindetails models.Login
	if err := c.BindJSON(&logindetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in proper format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	login, err := l.loginusecase.Login(logindetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "logigged in successfully", login, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *LoginHNadler) Signup(c *gin.Context) {
	var usersignup models.Users

	if err := c.BindJSON(&usersignup); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in proper format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err := i.loginusecase.Signup(usersignup)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "signup  successfully completed", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

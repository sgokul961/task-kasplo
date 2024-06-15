package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/pkg/demodb"
	"main.go/pkg/models"
	"main.go/pkg/usecase/usecaseinterface"
	"main.go/pkg/utils/response"
)

type LoginHNadler struct {
	loginusecase usecaseinterface.UseCase
}

func NewHandler(usecase usecaseinterface.UseCase) *LoginHNadler {
	return &LoginHNadler{
		loginusecase: usecase,
	}
}
func (l *LoginHNadler) Login(c *gin.Context) {

	var logindetails models.Users

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
	res, err := i.loginusecase.Signup(usersignup)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "signup  successfully completed", res, nil)
	c.JSON(http.StatusOK, successRes)
}
func (t *LoginHNadler) MakeToDO(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		errRes := response.ClientResponse(http.StatusUnauthorized, "Unauthorized access", nil, "User ID not found in context")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	// Perform type assertion to convert userID to int
	userIDInt, ok := userID.(int)
	if !ok {
		errRes := response.ClientResponse(http.StatusUnauthorized, "Unauthorized access", nil, "Invalid user ID format")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	var todo demodb.ToDo
	fmt.Println("called first")
	if err := c.BindJSON(&todo); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Check enterd details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	todo.UserID = userIDInt

	res, err := t.loginusecase.AddToDo(todo)
	fmt.Println("called")
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, " Check entered details ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "to-do created", res, nil)
	c.JSON(http.StatusOK, successRes)

}

func (t *LoginHNadler) UpdateToDo(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check path parameater", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		errRes := response.ClientResponse(http.StatusUnauthorized, "Unauthorized access", nil, "User ID not found in context")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	// Type assertion to convert userID to int
	userIDInt, ok := userID.(int)
	if !ok {
		errRes := response.ClientResponse(http.StatusInternalServerError, "Error converting user ID", nil, "User ID conversion error")
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	var update demodb.UpdateToDo

	if err := c.ShouldBindJSON(&update); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Invalid input data", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}
	updatedToDo, err := t.loginusecase.UpdateToDO(update, id, userIDInt)

	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "Failed to update ToDo", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "ToDo updated successfully", updatedToDo, nil)
	c.JSON(http.StatusOK, successRes)
}
func (t *LoginHNadler) DeleteToDo(c *gin.Context) {
	// Get the ToDo ID from the path parameters
	idStr := c.Query("id")
	if idStr == "" {
		errRes := response.ClientResponse(http.StatusBadRequest, "Invalid ToDo ID", nil, "ToDo ID is required")
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Invalid ToDo ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	// Get the user ID from the context (set by the AuthMiddleware)
	userID, exists := c.Get("userID")
	if !exists {
		errRes := response.ClientResponse(http.StatusUnauthorized, "Unauthorized access", nil, "User ID not found in context")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	// Type assertion to convert userID to int
	userIDInt, ok := userID.(int)
	if !ok {
		errRes := response.ClientResponse(http.StatusInternalServerError, "Error converting user ID", nil, "User ID conversion error")
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	// Call the use case to delete the ToDo
	result, err := t.loginusecase.DeleteToDO(id, userIDInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "Failed to delete ToDo", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	// Return the successful response
	successRes := response.ClientResponse(http.StatusOK, "TO-DO deleted successfully", result, nil)
	c.JSON(http.StatusOK, successRes)
}

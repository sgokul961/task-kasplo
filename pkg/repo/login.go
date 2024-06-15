package repo

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"main.go/pkg/demodb"
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
func (l *LoginRepo) Login(logindetails models.Users) (models.Users, error) {
	fmt.Println("logg", logindetails.Email)
	var login models.Users

	if err := l.DB.Raw(`SELECT * FROM users WHERE email=$1 `, logindetails.Email).Scan(&login).Error; err != nil {
		return models.Users{}, err
	}
	fmt.Println("user_ aher ", login)
	return models.Users{
		UserID:   login.UserID,
		Name:     login.Name,
		Email:    login.Email,
		Password: login.Password,
	}, nil

}
func (l *LoginRepo) Signup(signup models.Users) (models.SignupRes, error) {
	var userdetails models.SignupRes
	err := l.DB.Exec(`INSERT INTO users (name,email,password) values (?,?,?)`, signup.Name, signup.Email, signup.Password).Error

	if err != nil {
		return models.SignupRes{}, err
	}

	err = l.DB.Model(&models.Users{}).Select("user_id, email, name").Where("email = ? AND name =?", signup.Email, signup.Name).Scan(&userdetails).Error
	if err != nil {
		return models.SignupRes{}, err
	}
	return userdetails, nil

}
func (l *LoginRepo) CheckUserExistance(email string) bool {
	var count int

	query := `SELECT COUNT(*) FROM users WHERE email = ?`
	if err := l.DB.Raw(query, email).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}
func (l *LoginRepo) FetchUserEmail(id int) (string, error) {
	var email string
	query := `SELECT email FROM users WHERE user_id = ?`
	if err := l.DB.Raw(query, id).Scan(&email).Error; err != nil {
		return "", err
	}
	return email, nil

}
func (l *LoginRepo) CheckUserIdExist(id int) bool {
	var count int64
	query := `SELECT COUNT(*) FROM users WHERE user_id = ?`
	if err := l.DB.Raw(query, id).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (l *LoginRepo) AddToDo(content demodb.ToDo) (demodb.ToDo, error) {
	// Use the values from the content parameter
	query := `
        INSERT INTO to_dos (title, description, due_date, is_completed, created_at, updated_at, user_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id
    `

	err := l.DB.Raw(query, content.Title, content.Description, content.DueDate, content.IsCompleted, time.Now(), time.Now(), content.UserID).Scan(&content.ID).Error
	if err != nil {
		return demodb.ToDo{}, err
	}
	return content, nil
}
func (l *LoginRepo) UpdateToDo(content demodb.UpdateToDo) (demodb.UpdateToDo, error) {
	query := `
        UPDATE to_dos 
        SET title = $1, description = $2, due_date = $3, is_completed = $4, updated_at = $5
        WHERE id = $6
        RETURNING id, title, description, due_date, is_completed, created_at, updated_at, user_id
    `
	err := l.DB.Raw(query, content.Title, content.Description, content.DueDate, content.IsCompleted, time.Now(), content.ID).
		Scan(&content).Error
	if err != nil {
		return demodb.UpdateToDo{}, err
	}
	return content, nil
}
func (l *LoginRepo) DeleteToDo(id int) (string, error) {
	query := `DELETE FROM to_dos WHERE id = ?`

	if err := l.DB.Exec(query, id).Error; err != nil {
		return "", err
	}

	return "ToDo deleted successfully", nil
}

func (i *LoginRepo) CheToDOExist(id int) bool {
	var count int64

	query := `SELECT id FROM to_dos WHERE id =?`
	if err := i.DB.Raw(query, id).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0

}
func (i *LoginRepo) GetToDoByID(id int) (demodb.ToDo, error) {
	var todo demodb.ToDo
	query := `SELECT * FROM to_dos WHERE id = ?`
	if err := i.DB.Raw(query, id).Scan(&todo).Error; err != nil {
		return demodb.ToDo{}, err
	}
	return todo, nil
}

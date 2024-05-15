package utils

import "main.go/pkg/models"

type Admin struct {
	Name  string `json:"name" gorm:"validate:required"`
	Email string `json:"email" gorm:"validate:required"`
}
type TokenLOgin struct {
	Admin        models.Login
	Token        string
	RefreshToken string
}

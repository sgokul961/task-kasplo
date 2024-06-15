package utils

type LoginResponse struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}
type TokenLOgin struct {
	LoginResponse LoginResponse
	Token         string
}

package models

type Login struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password" validate:"min=8,max=20"`
}

type Users struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=8,max==20"`
}

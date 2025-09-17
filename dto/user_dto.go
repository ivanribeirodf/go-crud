package dto

type CreateUserDTO struct {
	Name   string `json:"name" binding:"required,min=3"`
	Email  string `json:"email" binding:"required,email"`
	Passwd string `json:"password" binding:"required,min=6"`
}

type LoginDTO struct {
	Email  string `json:"email" binding:"required,email"`
	Passwd string `json:"password" binding:"required"`
}

type UpdateUserDTO struct {
	Name  string `json:"name" binding:"omitempty,min=3"`
	Email string `json:"email" binding:"omitempty,email"`
}

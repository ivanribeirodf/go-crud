package dto

type CreateUserDTO struct {
    Name  string `json:"name" binding:"required,min=3"`
    Email string `json:"email" binding:"required,email"`
}

type UpdateUserDTO struct {
    Name  string `json:"name" binding:"omitempty,min=3"`
    Email string `json:"email" binding:"omitempty,email"`
}
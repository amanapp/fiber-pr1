package user

type CreateUserDTO struct {
	Name  string `json:"name" validate:"required" example:"Aman"`
	Email string `json:"email" validate:"required,email" example:"aman@gmail.com"`
}

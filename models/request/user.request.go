package request

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Age      uint8  `json:"age" validate:"number"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}

type UserUpdateRequest struct {
	Name    string `json:"name"`
	Age     uint8  `json:"age" validate:"number"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

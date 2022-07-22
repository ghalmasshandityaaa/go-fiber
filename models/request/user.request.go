package request

type UserCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Age     uint8  `json:"age" validate:"number"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type UserUpdateRequest struct {
	Name    string `json:"name"`
	Age     uint8  `json:"age" validate:"number"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

package request

type UserCreateRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     uint8  `json:"age"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

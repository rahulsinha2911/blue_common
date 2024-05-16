package requests

type UserRegisterationRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Sex   string `json:"sex"`
	Age   int    `json:"age"`
}

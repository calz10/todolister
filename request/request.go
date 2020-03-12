package request

type UserRequestBody struct {
	Email    string `json:"email" validate:"required,email,required_without=EmailVerified"`
	Password string `json:"password" validate:"min=8,max=32,required"`
	Username string `json:"username" validate:"required"`
}

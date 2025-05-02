package auth

type signUpReq struct {
	Username string `json:"username" validate:"required,min=3,max=30,alphanum"`
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email"    validate:"required,email"`
}

type signInReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

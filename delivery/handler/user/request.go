package user

type loginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

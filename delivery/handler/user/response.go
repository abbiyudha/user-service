package user

type UserResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginUserResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

type CreateUserResponse struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
}

package user

type UserResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UnAuthorizeResponse struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
}

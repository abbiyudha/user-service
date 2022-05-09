package user

import (
	"CodingTestUser/delivery/helper"
	"CodingTestUser/delivery/middleware"
	"CodingTestUser/entities"
	"CodingTestUser/usecase/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	userUseCase user.UserUsecaseInferface
}

func NewUserHandler(userUsecase user.UserUsecaseInferface) *UserHandler {
	return &UserHandler{
		userUseCase: userUsecase,
	}
}

func (uh *UserHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userParam entities.User
		_, role := middleware.ExtractToken(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, UnAuthorizeResponse{
				Status:   "Failed",
				Messages: "Unauthorized",
			})
		}

		c.Bind(&userParam)

		userParam.Role = "User"
		err := uh.userUseCase.CreateUser(userParam)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success create user"))
	}
}

func (uh *UserHandler) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		type loginData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		var login loginData

		err := c.Bind(&login)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data"))
		}

		token, errorLogin := uh.userUseCase.LoginUser(login.Email, login.Password)
		if errorLogin != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errorLogin.Error()))
		}
		responseToken := map[string]interface{}{
			"token": token,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success login", responseToken))
	}
}

func (uh *UserHandler) GetUserById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := middleware.ExtractToken(c)

		user, err := uh.userUseCase.GetUserById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("user not found"))
		}
		var UserResponse = UserResponse{
			Id:    user.IdString,
			Name:  user.Name,
			Email: user.Email,
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("Success get user by id", UserResponse))
	}

}

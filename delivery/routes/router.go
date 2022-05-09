package routes

import (
	"CodingTestUser/delivery/handler/user"
	"CodingTestUser/delivery/middleware"
	"github.com/labstack/echo/v4"
)

func UserPath(e *echo.Echo, uh *user.UserHandler) {
	e.POST("create/user", uh.CreateUser(), middleware.JWTMiddleware())
	e.POST("/user/login", uh.LoginUserHandler())
	e.GET("/user", uh.GetUserById(), middleware.JWTMiddleware())
}

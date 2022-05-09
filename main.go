package main

import (
	"CodingTestUser/configs"
	"CodingTestUser/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"

	_userHandler "CodingTestUser/delivery/handler/user"
	_userRepo "CodingTestUser/repository/user"
	_userUsecase "CodingTestUser/usecase/user"

	"CodingTestUser/delivery/routes"
)

func main() {
	config := configs.GetConfig()
	db, _ := utils.Connect(config)

	userRepo := _userRepo.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUsecase)

	e := echo.New()
	routes.UserPath(e, userHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}

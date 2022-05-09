package user

import (
	"CodingTestUser/delivery/middleware"
	"CodingTestUser/entities"
	"CodingTestUser/repository/user"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseInferface interface {
	CreateUser(UserParam entities.User) error
	LoginUser(email string, password string) (string, error)
	GetUserById(id string) (entities.User, error)
}

type UserUseCase struct {
	UserRepository user.UserRepositoryInferface
}

func NewUserUsecase(userRepo user.UserRepositoryInferface) UserUsecaseInferface {
	return &UserUseCase{
		UserRepository: userRepo,
	}
}

func (uuc *UserUseCase) CreateUser(UserParam entities.User) error {
	err := uuc.UserRepository.CreateUser(UserParam)
	return err
}

func (uuc *UserUseCase) LoginUser(email string, password string) (string, error) {
	userData, err := uuc.UserRepository.LoginUser(email)
	if err != nil {
		return "", errors.New("Email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if err != nil {
		return "", errors.New("Wrong Password")
	}

	userData.IdString = userData.Id.Hex()
	token, _ := middleware.CreateToken(userData.IdString, userData.Name)

	return token, err
}

func (uuc *UserUseCase) GetUserById(id string) (entities.User, error) {
	UserParam, err := uuc.UserRepository.GetUserById(id)

	UserParam.IdString = UserParam.Id.Hex()
	return UserParam, err
}

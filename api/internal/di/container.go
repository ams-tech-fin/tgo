package di

import (
	userUseCase "tgo/api/internal/application/usecase/user"
	"tgo/api/internal/modules/crypto"
	"tgo/api/internal/modules/database"
	userController "tgo/api/internal/modules/http/controllers"
	userRepository "tgo/api/internal/repositories/user"
	userService "tgo/api/internal/services/user"

	"github.com/jmoiron/sqlx"
)

type Container struct {
	UserController *userController.UserController
	DB             *sqlx.DB
}

func NewContainer() (*Container, error) {

	db, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}

	cryptoAdapter := &crypto.Argon2Adapter{}

	userRepoInit := userRepository.NewUserRepository(db)
	createUserUseCaseInit := userUseCase.NewCreateUserUseCase(userRepoInit)
	getUserByIdUseCaseInit := userUseCase.NewGetUserByIdUseCase(userRepoInit)
	getAllUsersUseCaseInit := userUseCase.NewGetAllUsersUseCase(userRepoInit)
	userServiceInit := userService.NewUserService(cryptoAdapter, createUserUseCaseInit, getUserByIdUseCaseInit, getAllUsersUseCaseInit)
	userControllerInit := userController.NewUserController(userServiceInit)

	return &Container{
		UserController: userControllerInit,
		DB:             db,
	}, nil
}

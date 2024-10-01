package di

import (
	"context"
	"log"
	userUseCase "tgo/api/internal/application/usecase/user"
	"tgo/api/internal/modules/cache"
	"tgo/api/internal/modules/crypto"
	"tgo/api/internal/modules/database"
	controller "tgo/api/internal/modules/http/controllers"
	userRepository "tgo/api/internal/repositories/user"
	userService "tgo/api/internal/services/user"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type Container struct {
	UserController  *controller.UserController
	CacheController *controller.CacheController
	DB              *sqlx.DB
	Cache           *redis.Client
}

func NewContainer() (*Container, error) {

	db, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}

	cryptoAdapter := &crypto.Argon2Adapter{}

	cache.SetupRedis()
	ctx := context.Background()
	_, err = cache.PingRedis(ctx)
	if err != nil {
		log.Fatalf("Erro ao conectar no Redis: %v", err)
	}

	userRepoInit := userRepository.NewUserRepository(db)
	createUserUseCaseInit := userUseCase.NewCreateUserUseCase(userRepoInit)
	getUserByIdUseCaseInit := userUseCase.NewGetUserByIdUseCase(userRepoInit)
	getAllUsersUseCaseInit := userUseCase.NewGetAllUsersUseCase(userRepoInit)
	getAuthUserUseCaseInit := userUseCase.NewGetAuthUserUseCase(userRepoInit, cryptoAdapter)
	userServiceInit := userService.NewUserService(
		cryptoAdapter,
		createUserUseCaseInit,
		getUserByIdUseCaseInit,
		getAllUsersUseCaseInit,
		getAuthUserUseCaseInit,
		cache.Rdb,
	)
	userControllerInit := controller.NewUserController(userServiceInit)

	cacheControllerInit := controller.NewCacheController(cache.Rdb)
	return &Container{
		UserController:  userControllerInit,
		CacheController: cacheControllerInit,
		DB:              db,
	}, nil
}

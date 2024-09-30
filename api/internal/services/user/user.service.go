package service

import (
	"context"
	userUseCase "tgo/api/internal/application/usecase/user"
	"tgo/api/internal/modules/crypto"
	repository "tgo/api/internal/repositories/user"

	"github.com/google/uuid"
)

type UserService struct {
	cryptoAdapter      *crypto.Argon2Adapter
	createUserUseCase  userUseCase.ICreateUserUseCase
	getUserByIdUseCase userUseCase.IGetUserByIdUseCase
	getAllUsersUseCase userUseCase.IGetAllUsersUseCase
}

func NewUserService(cryptoAdapter *crypto.Argon2Adapter, createUserUseCase userUseCase.ICreateUserUseCase, getUserByIdUseCase userUseCase.IGetUserByIdUseCase, getAllUsersUseCase userUseCase.IGetAllUsersUseCase) *UserService {
	return &UserService{cryptoAdapter: cryptoAdapter, createUserUseCase: createUserUseCase, getUserByIdUseCase: getUserByIdUseCase, getAllUsersUseCase: getAllUsersUseCase}
}

func (s *UserService) CreateUser(ctx context.Context, name string, email string, password string) (repository.User, error) {
	hash, _ := s.cryptoAdapter.Hash(password)
	return s.createUserUseCase.ExecuteNewCreateUserUseCase(ctx, name, email, hash)
}

func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (repository.User, error) {
	return s.getUserByIdUseCase.ExecuteGetUserByIdUseCase(ctx, id)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]repository.User, error) {
	return s.getAllUsersUseCase.ExecuteGetAllUsersUseCase(ctx)
}

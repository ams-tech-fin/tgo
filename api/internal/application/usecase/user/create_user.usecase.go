package usecase

import (
	"context"
	repository "tgo/api/internal/repositories/user"
)

type ICreateUserUseCase interface {
	ExecuteNewCreateUserUseCase(ctx context.Context, name string, email string, password string) (repository.User, error)
}

type CreateUserUseCase struct {
	repo *repository.UserRepository
}

func NewCreateUserUseCase(repo *repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

func (uc *CreateUserUseCase) ExecuteNewCreateUserUseCase(ctx context.Context, name string, email string, password string) (repository.User, error) {
	return uc.repo.CreateUser(ctx, name, email, password)
}

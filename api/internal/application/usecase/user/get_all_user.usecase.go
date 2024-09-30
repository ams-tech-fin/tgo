package usecase

import (
	"context"
	repository "tgo/api/internal/repositories/user"
)

type IGetAllUsersUseCase interface {
	ExecuteGetAllUsersUseCase(ctx context.Context) ([]repository.User, error)
}

type GetAllUsersUseCase struct {
	repo *repository.UserRepository
}

func NewGetAllUsersUseCase(repo *repository.UserRepository) *GetAllUsersUseCase {
	return &GetAllUsersUseCase{repo: repo}
}

func (uc *GetAllUsersUseCase) ExecuteGetAllUsersUseCase(ctx context.Context) ([]repository.User, error) {
	return uc.repo.GetAllUsers(ctx)
}

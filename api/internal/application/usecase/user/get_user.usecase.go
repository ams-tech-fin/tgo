package usecase

import (
	"context"
	repository "tgo/api/internal/repositories/user"

	"github.com/google/uuid"
)

type IGetUserByIdUseCase interface {
	ExecuteGetUserByIdUseCase(ctx context.Context, id uuid.UUID) (repository.User, error)
}

type GetUserByIdUseCase struct {
	repo *repository.UserRepository
}

func NewGetUserByIdUseCase(repo *repository.UserRepository) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{repo: repo}
}

func (uc *GetUserByIdUseCase) ExecuteGetUserByIdUseCase(ctx context.Context, id uuid.UUID) (repository.User, error) {
	return uc.repo.GetUserByID(ctx, id)
}

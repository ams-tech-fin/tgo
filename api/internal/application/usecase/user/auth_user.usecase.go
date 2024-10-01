package usecase

import (
	"context"
	"tgo/api/internal/modules/crypto"
	repository "tgo/api/internal/repositories/user"
)

type IGetAuthUserUseCase interface {
	ExecuteGetAuthUserUseCase(ctx context.Context, email string, password string) (repository.User, error)
}

type GetAuthUserUseCase struct {
	repo          *repository.UserRepository
	cryptoAdapter *crypto.Argon2Adapter
}

func NewGetAuthUserUseCase(repo *repository.UserRepository, cryptoAdapter *crypto.Argon2Adapter) *GetAuthUserUseCase {
	return &GetAuthUserUseCase{repo: repo, cryptoAdapter: cryptoAdapter}
}

func (uc *GetAuthUserUseCase) ExecuteGetAuthUserUseCase(ctx context.Context, email string, password string) (repository.User, error) {

	user, _ := uc.repo.GetUserByEmail(ctx, email)
	isValid, err2 := uc.cryptoAdapter.Verify(user.Password, password)

	if isValid {
		return user, err2
	} else {
		return user, err2
	}

}

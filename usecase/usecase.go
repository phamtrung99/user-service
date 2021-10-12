package usecase

import (
	"github.com/phamtrung99/user-service/repository"
	"github.com/phamtrung99/user-service/usecase/user"
	"github.com/phamtrung99/user-service/usecase/userfavorite"
)

type UseCase struct {
	User      user.IUsecase
	UserFavor userfavorite.IUsecase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		User:      user.New(repo),
		UserFavor: userfavorite.New(repo),
	}
}

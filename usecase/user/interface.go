package user

import (
	"context"

	"github.com/phamtrung99/user-service/model"
)

// IUsecase .
type IUsecase interface {
	Update(ctx context.Context, req UpdateInfoRequest) (*model.User, error)
	ChangePassword(ctx context.Context, req ChangePasswRequest) error
	GetMe(ctx context.Context) (*model.User, error)
}

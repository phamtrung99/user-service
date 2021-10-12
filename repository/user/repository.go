package user

import (
	"context"

	"github.com/phamtrung99/user-service/model"
)

// Repository .
type Repository interface {
	Delete(ctx context.Context, id int64) error
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	CheckEmailExist(ctx context.Context, email string) bool
	GetByID(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	GetNextIDIncrement(ctx context.Context) string
	UpdateUserPassword(ctx context.Context, passwHash string, id int64) error
}

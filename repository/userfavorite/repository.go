package userfavorite

import (
	"context"

	"github.com/phamtrung99/user-service/model"
)

// Repository .
type Repository interface {
	Create(ctx context.Context, userfavor *model.UserFavorite) (*model.UserFavorite, error)
	GetListMovieIDByUserID(ctx context.Context, userID int64) ([]int64, error)
	GetByIDMovieAndIDUser(ctx context.Context, userID int64, movieID int64) (*model.UserFavorite, error)
}

package userfavorite

import (
	"context"

	"github.com/phamtrung99/user-service/model"
	moviemodel"github.com/phamtrung99/movie-service/model"
)

// IUsecase .
type IUsecase interface {
	AddMovieToFavorite(ctx context.Context, req AddFavoriteRequest) (*model.UserFavorite, error)
	GetFavoriteMovie(ctx context.Context, req UserFavorRequest) (*moviemodel.MovieResult, error)
}

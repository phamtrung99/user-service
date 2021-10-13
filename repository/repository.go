package repository

import (
	"context"

	"github.com/phamtrung99/movie-service/repository/movie"
	"github.com/phamtrung99/user-service/repository/user"
	"github.com/phamtrung99/user-service/repository/userfavorite"
	"gorm.io/gorm"
)

type Repository struct {
	User      user.Repository
	UserFavor userfavorite.Repository
	Movie     movie.Repository
}

func New(
	getSQLClient func(ctx context.Context) *gorm.DB,
	// getRedisClient func(ctx context.Context) *redis.Client,
) *Repository {
	return &Repository{
		User:      user.NewPGRepository(getSQLClient),
		UserFavor: userfavorite.NewPGRepository(getSQLClient),
		Movie:     movie.NewPGRepository(getSQLClient),
	}
}

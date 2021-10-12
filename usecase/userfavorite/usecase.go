package userfavorite

import (
	"github.com/phamtrung99/user-service/repository"
	"github.com/phamtrung99/movie-service/repository/movie"
	"github.com/phamtrung99/user-service/repository/userfavorite"
)

type Usecase struct {
	userFavorRepo userfavorite.Repository
	movieRepo     movie.Repository
}

// New .
func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		userFavorRepo: repo.UserFavor,
		movieRepo:     repo.Movie,
	}
}

package userfavorite

import (
	"context"

	"github.com/phamtrung99/gopkg/middleware"
	"github.com/phamtrung99/user-service/model"
	"github.com/phamtrung99/user-service/util/myerror"
)

type AddFavoriteRequest struct {
	MovieID int64 `json:"movie_id"`
}

func (u *Usecase) AddMovieToFavorite(ctx context.Context, req AddFavoriteRequest) (*model.UserFavorite, error) {

	//Get current userId from Token.
	claim := middleware.GetClaim(ctx)
	userID := claim.UserID

	userFavor := &model.UserFavorite{}
	userFavor.MovieID = req.MovieID
	userFavor.UserID = userID

	userFavorTemp, err := u.userFavorRepo.GetByIDMovieAndIDUser(ctx, userFavor.UserID, userFavor.MovieID)

	if err != nil {
		return nil, err
	}

	if userFavorTemp.ID != 0 {
		return nil, myerror.ErrFavorMovieExist(nil)
	}

	result, err := u.userFavorRepo.Create(ctx, userFavor)
	if err != nil {
		return &model.UserFavorite{}, err
	}

	return result, nil
}

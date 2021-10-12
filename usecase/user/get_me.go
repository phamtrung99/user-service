package user

import (
	"context"

	"github.com/phamtrung99/gopkg/middleware"
	"github.com/phamtrung99/user-service/model"
)

func (u *Usecase) GetMe(ctx context.Context) (*model.User, error) {
	//Get current userId from Token.
	claim := middleware.GetClaim(ctx)
	userID := claim.UserID

	user, err := u.userRepo.GetByID(ctx, userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

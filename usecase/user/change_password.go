package user

import (
	"context"

	"github.com/phamtrung99/gopkg/middleware"
	"github.com/phamtrung99/auth-service/package/auth"
	checkform "github.com/phamtrung99/movie-service/package/checkForm"
	"github.com/phamtrung99/user-service/util/myerror"
	authmyerror"github.com/phamtrung99/auth-service/util/myerror"
)

type ChangePasswRequest struct {
	OldPassword   string `json:"old_password"`
	NewPassword   string `json:"new_password"`
	RenewPassword string `json:"renew_password"`
}

func (u *Usecase) ChangePassword(ctx context.Context, req ChangePasswRequest) error {

	if req.NewPassword != req.RenewPassword {
		return myerror.ErrPwdMatching(nil)
	}

	isPass, newPwd := checkform.CheckFormatValue("new_password", req.NewPassword)
	if !isPass {
		return myerror.ErrNewPwdFormat(nil)
	}
	isPass, oldPwd := checkform.CheckFormatValue("old_password", req.OldPassword)
	if !isPass {
		return myerror.ErrOldPwdFormat(nil)
	}

	//Get current userId from Token.
	claim := middleware.GetClaim(ctx)
	userID := claim.UserID

	//Get current user info from userID
	user, err := u.userRepo.GetByID(ctx, userID)

	if err != nil {
		return err
	}

	//Check old password is true.
	isPassTrue := auth.VerifyPassword(oldPwd, user.Password)

	if !isPassTrue {
		return authmyerror.ErrInvalid(nil)
	}

	passHash, err := auth.HashPassword(newPwd)

	if err != nil {
		return authmyerror.ErrHashPassword(err)
	}

	user.Password = passHash

	_, err = u.userRepo.Update(ctx, user)

	if err != nil {
		return err
	}

	return nil
}
